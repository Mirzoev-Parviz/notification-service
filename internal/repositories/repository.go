package repositories

import "database/sql"

type Repository struct {
	Event
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Event: NewEventRepository(db),
	}
}
