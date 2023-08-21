package handler

import (
	"context"
	"fmt"
	"github.com/nduyphuong/gorya/internal/store"
	"github.com/nduyphuong/gorya/internal/worker"
	"github.com/nduyphuong/gorya/pkg/api/service/v1alpha1"
	"github.com/nduyphuong/gorya/pkg/aws"
	"github.com/nduyphuong/gorya/pkg/timezone"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func ScheduleTaskV1alpha1(ctx context.Context, awsClient aws.Interface, store store.Interface,
	taskProcessor worker.Interface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("hello")
		var err error
		policies, err := store.ListPolicy()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, policy := range *policies {
			schedule, err := store.GetSchedule(policy.Name)
			if err == gorm.ErrRecordNotFound {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Println(schedule.TimeZone)
			location, err := time.LoadLocation(schedule.TimeZone)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			nowInTargetLocation := time.Now().In(location)
			day, hour := timezone.ConvertTimeToIndex(nowInTargetLocation)
			var arr []int
			for _, v := range schedule.Schedule.Data().NdArray {
				arr = append(arr, v...)
			}
			matrixSize := schedule.Schedule.Data().Shape[0] * schedule.Schedule.Data().Shape[1]
			prevIdx := getPreviousIdx(day*24+hour, matrixSize)
			now := arr[day*24+hour]
			prev := arr[prevIdx]
			if now != prev {
				for _, tag := range policy.Tags {
					for k, v := range tag {
						for _, project := range policy.Projects {
							e := worker.QueueElem{
								RequestURI: v1alpha1.GoryaTaskChangeStageProcedure,
								Project:    project,
								TagKey:     k,
								TagValue:   v,
								Action:     now,
							}
							taskProcessor.Dispatch(ctx, &e)
						}
					}
				}
			}
		}
	}
}

func getPreviousIdx(idx int, matrixSize int) int {
	if idx == 0 {
		return matrixSize - 1
	}
	return idx - 1
}
