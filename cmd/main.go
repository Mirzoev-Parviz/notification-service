package main

import (
	"context"
	"fmt"
	"log"
	"notification_service"
	"notification_service/internal/db"
	"notification_service/internal/handlers"
	"notification_service/internal/repositories"
	"notification_service/internal/services"
	"notification_service/internal/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	utils.ReadSettings()
	db.StartDBConnection()
	_db := db.GetDBConn()
	server := new(notification_service.Server)

	repositories := repositories.NewRepository(_db)
	services := services.NewService(repositories)
	handlers := handlers.NewHandler(services)

	go func() {
		if err := server.Run("localhost:8080", handlers.InitRoutes()); err != nil {
			log.Fatal(err.Error())
		}
	}()

	fmt.Println("App started on port 8080")

	go services.Worker(time.Second * 5)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Println(err.Error())
	}
}
