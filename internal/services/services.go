package services

import "notification_service/internal/repositories"

type Service struct {
	Job
	Event
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		Job:   NewJobService(repo),
		Event: NewEventRepository(repo),
	}
}
