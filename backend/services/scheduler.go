package services

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	scheduler *gocron.Scheduler
}

func NewScheduler() *Scheduler {
	s := gocron.NewScheduler(time.UTC)
	return &Scheduler{scheduler: s}
}

func (s *Scheduler) StartScheduler() {
	log.Println("Starting the scheduler...")
	s.scheduler.StartAsync()
}

func (s *Scheduler) AddJob(interval time.Duration, jobFunc func()) {
	_, err := s.scheduler.Every(interval).Do(jobFunc)
	if err != nil {
		log.Fatalf("Failed to schedule job: %v", err)
	}
}

func PerformPeriodicTask(db *sql.DB) {
	statuses, err := GetServiceStatuses(db)
	if err != nil {
		log.Printf("Failed to fetch service statuses: %v\n", err)
		return
	}

	for _, status := range statuses {
		log.Printf("Service: %s, Status: %s, Uptime: %s\n", status.Name, status.Status, status.Uptime_duration)
	}
}
