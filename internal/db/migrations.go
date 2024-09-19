package db

import (
	"database/sql"
	"fmt"
)

func runMigrations(db *sql.DB) error {
	createTableQuery := `  
	CREATE TABLE IF NOT EXISTS events (  
		id SERIAL PRIMARY KEY,  
		order_type VARCHAR(50),  
		session_id VARCHAR(50),  
		card VARCHAR(50),  
		event_date TIMESTAMP,  
		website_url VARCHAR(255),  
		is_delivered BOOLEAN DEFAULT FALSE  
	);  
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	return nil
}
