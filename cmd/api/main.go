package main

import (
	"fmt"
	"log"

	"mypos-backend/internal/config"
	"mypos-backend/internal/database"
	"mypos-backend/internal/handler"
	"mypos-backend/internal/repository"
	"mypos-backend/internal/router"
	"mypos-backend/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := router.SetupRoutes(userHandler)

	serverAddr := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	log.Printf("Starting MyPOS backend server (Gin) on http://%s ...\n", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
