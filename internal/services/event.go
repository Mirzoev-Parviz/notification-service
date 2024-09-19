package services

import (
	"notification_service/internal/models"
	"notification_service/internal/repositories"
)

type Event interface {
	SaveEvent(event models.Event) error
}

type EventServices struct {
	repo *repositories.Repository
}

func NewEventRepository(repo *repositories.Repository) *EventServices {
	return &EventServices{repo: repo}
}

func (e *EventServices) SaveEvent(event models.Event) error {
	return e.repo.SaveEvent(event)
}
