package main

import (
	"log"
	"user-service/internal/pkg/config"
	"user-service/internal/pkg/postgres"
	RunService "user-service/internal/pkg/service"
	"user-service/internal/repository"
	"user-service/internal/service"
)

func main() {
	cfg, err := config.LOAD("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	db, err := postgres.InitDB(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	queries := postgres.QueriesDB(db)
	userRepo := repository.NewUserREPO(db, queries)
	service := service.NewService(userRepo)
	runService := RunService.NewRunService(&service)

	log.Printf("User Service running on :%d port", cfg.ServicePort)
	if err := runService.Run(cfg); err != nil {
		log.Fatal(err)
	}

}
