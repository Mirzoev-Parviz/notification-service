package services

import (
	"fmt"
	"log"
	"notification_service/internal/models"
	"notification_service/internal/repositories"
	"time"
)

type Job interface {
	Worker(interval time.Duration)
}

type JobService struct {
	repo *repositories.Repository
}

func NewJobService(repo *repositories.Repository) *JobService {
	return &JobService{repo: repo}
}

func (j *JobService) Worker(interval time.Duration) {
	for {
		time.Sleep(interval)

		events, err := j.repo.GetUndeliveredEvents()
		if err != nil {
			log.Println("Error fetching undelivered events:", err)
			continue
		}

		for _, event := range events {
			j.NotifyUser(event)
			if err := j.repo.UpdateEventToDelivered(event); err != nil {
				log.Println("Error updating event to delivered:", err)
			}
		}

	}
}

func (j *JobService) NotifyUser(event models.Event) {
	fmt.Printf("Notification: Card %s used for %s on %s at %s\n",
		event.Card, event.OrderType, event.WebsiteUrl, event.EventDate)
}
