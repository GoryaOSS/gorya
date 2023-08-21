package api

import (
	"context"
	"github.com/nduyphuong/gorya/internal/api/config"
	"github.com/nduyphuong/gorya/internal/api/handler"
	httputil "github.com/nduyphuong/gorya/internal/http"
	"github.com/nduyphuong/gorya/internal/logging"
	"github.com/nduyphuong/gorya/internal/os"
	queueOptions "github.com/nduyphuong/gorya/internal/queue/options"
	"github.com/nduyphuong/gorya/internal/store"
	"github.com/nduyphuong/gorya/internal/version"
	"github.com/nduyphuong/gorya/internal/worker"
	svcv1alpha1 "github.com/nduyphuong/gorya/pkg/api/service/v1alpha1"
	"github.com/nduyphuong/gorya/pkg/aws"
	"github.com/pkg/errors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net"
	"net/http"
	goos "os"
	"time"
)

type Server interface {
	Serve(ctx context.Context, l net.Listener) error
}

type server struct {
	cfg           config.ServerConfig
	sc            store.Interface
	aws           aws.Interface
	taskProcessor worker.Interface
}

func NewServer(cfg config.ServerConfig) (Server, error) {
	return &server{
		cfg: cfg,
	}, nil
}

func (s *server) Serve(ctx context.Context, l net.Listener) error {
	var err error
	errCh := make(chan error)
	log := logging.LoggerFromContext(ctx)
	log.Infof("Server is listening on %q", l.Addr().String())
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	s.sc, err = store.GetOnce()
	if err != nil {
		return err
	}
	awsRegion := os.GetEnv("AWS_REGION", "ap-southeast-1")
	s.aws, err = aws.New(ctx, awsRegion)
	if err != nil {
		return err
	}
	s.taskProcessor = worker.NewClient(worker.Options{
		QueueOpts: queueOptions.Options{
			Addr:          os.GetEnv("GORYA_REDIS_ADDR", "localhost:6379"),
			Name:          os.GetEnv("GORYA_QUEUE_NAME", "gorya"),
			FetchInterval: 2 * time.Second,
		},
	})
	path, svcHandler := svcv1alpha1.NewGoryaServiceHandler(ctx, s.sc, s)
	mux.Handle(path, svcHandler)
	mux.Handle("/ui", s.newDashboardRequestHandler())
	srv := &http.Server{
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: time.Minute,
	}
	go func() { errCh <- srv.Serve(l) }()
	select {
	case <-ctx.Done():
		log.Info("Gracefully stopping server...")
		time.Sleep(s.cfg.GracefulShutdownTimeout)
		return srv.Shutdown(context.Background())
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}

func (s *server) newDashboardRequestHandler() http.HandlerFunc {
	fs := http.FileServer(http.Dir(s.cfg.UIDirectory))
	return func(w http.ResponseWriter, req *http.Request) {
		path := s.cfg.UIDirectory + req.URL.Path
		info, err := goos.Stat(path)
		if goos.IsNotExist(err) || info.IsDir() {
			if w != nil {
				httputil.SetNoCacheHeaders(w)
				http.ServeFile(w, req, s.cfg.UIDirectory+"/index.html")
			}
		} else {
			fs.ServeHTTP(w, req)
		}
	}
}

func (s *server) GetTimeZone() http.Handler {
	return handler.GetTimeZoneV1Alpha1()
}

func (s *server) GetVersionInfo() http.Handler {
	return handler.GetVersionInfoV1Alpha1(version.GetVersion())
}

func (s *server) AddSchedule(ctx context.Context) http.Handler {
	return handler.AddScheduleV1Alpha1(ctx, s.sc)
}

func (s *server) GetSchedule(ctx context.Context) http.Handler {
	return handler.GetScheduleV1alpha1(ctx, s.sc)
}

func (s *server) ListSchedule(ctx context.Context) http.Handler {
	return handler.ListScheduleV1alpha1(ctx, s.sc)
}

func (s *server) DeleteSchedule(ctx context.Context) http.Handler {
	return handler.DeleteScheduleV1alpha1(ctx, s.sc)
}

func (s *server) AddPolicy(ctx context.Context) http.Handler {
	return handler.AddPolicyV1Alpha1(ctx, s.sc)
}

func (s *server) GetPolicy(ctx context.Context) http.Handler {
	return handler.GetPolicyV1Alpha1(ctx, s.sc)
}

func (s *server) ListPolicy(ctx context.Context) http.Handler {
	return handler.ListPolicyV1alpha1(ctx, s.sc)
}

func (s *server) DeletePolicy(ctx context.Context) http.Handler {
	return handler.DeletePolicyV1alpha1(ctx, s.sc)
}

func (s *server) ChangeState(ctx context.Context) http.Handler {
	return handler.ChangeStateV1alpha1(ctx, s.aws)
}

func (s *server) ScheduleTask(ctx context.Context) http.Handler {
	return handler.ScheduleTaskV1alpha1(ctx, s.aws, s.sc, s.taskProcessor)
}
