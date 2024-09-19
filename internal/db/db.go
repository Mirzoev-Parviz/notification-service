package db

import (
	"database/sql"
	"fmt"
	"log"
	"notification_service/internal/utils"

	_ "github.com/lib/pq"
)

var database *sql.DB

func ConnectDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		utils.AppSettings.DBSettings.Host, utils.AppSettings.DBSettings.Port,
		utils.AppSettings.DBSettings.User, utils.AppSettings.DBSettings.Password,
		utils.AppSettings.DBSettings.Database)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func StartDBConnection() {
	database = ConnectDB()
	err := runMigrations(database)
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}
}

func GetDBConn() *sql.DB {
	return database
}

func DisconnectDB(db *sql.DB) error {
	if err := db.Close(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
