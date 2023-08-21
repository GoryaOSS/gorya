package store

import (
	"github.com/nduyphuong/gorya/internal/models"
)

type Interface interface {
	SavePolicy(policy models.Policy) error
	GetPolicyByName(name string) (*models.Policy, error)
	GetPolicyBySchedule(name string) (*[]models.Policy, error)
	ListPolicy() (*[]models.Policy, error)
	DeletePolicy(name string) error
	SaveSchedule(schedule models.ScheduleModel) error
	GetSchedule(name string) (*models.ScheduleModel, error)
	ListSchedule() (*[]models.ScheduleModel, error)
	DeleteSchedule(name string) error
}
