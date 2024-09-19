package repositories

import (
	"database/sql"
	"notification_service/internal/models"
)

type Event interface {
	SaveEvent(event models.Event) error
	GetUndeliveredEvents() ([]models.Event, error)
	UpdateEventToDelivered(event models.Event) error
}

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (e *EventRepository) SaveEvent(event models.Event) error {
	query := `INSERT INTO events (order_type, session_id, card, event_date, website_url, is_delivered)  
			  VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := e.db.Exec(query, event.OrderType, event.SessionID, event.Card, event.EventDate, event.WebsiteUrl, event.IsDelivered)
	return err
}

func (e *EventRepository) GetUndeliveredEvents() ([]models.Event, error) {
	query := `SELECT order_type, session_id, card, event_date, website_url, is_delivered FROM events WHERE is_delivered = FALSE`

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.OrderType, &event.SessionID, &event.Card, &event.EventDate, &event.WebsiteUrl, &event.IsDelivered); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (e *EventRepository) UpdateEventToDelivered(event models.Event) error {
	query := `UPDATE events SET is_delivered = TRUE WHERE session_id = $1 AND card = $2`

	_, err := e.db.Exec(query, event.SessionID, event.Card)
	return err
}
