package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	scheduler *gocron.Scheduler
}

func NewScheduler() *Scheduler {
	location, _ := time.LoadLocation("Asia/Anadyr")
	return &Scheduler{scheduler: gocron.NewScheduler(location)}
}
func (S *Scheduler) ResetStatus(time string, job func()) {
	S.scheduler.Every(1).Day().At(time).Do(job)

}
func (S *Scheduler) Start() {
	S.scheduler.StartAsync()
}
