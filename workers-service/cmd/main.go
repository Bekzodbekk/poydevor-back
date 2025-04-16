package main

import (
	"log"
	"workers-service/internal/pkg/config"
	InitPostgres "workers-service/internal/pkg/postgres"
	RunService "workers-service/internal/pkg/service"
	"workers-service/internal/repository"
	"workers-service/internal/service"
)

func main() {
	cfg, err := config.LOAD("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the database connection
	db, err := InitPostgres.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	Queries := InitPostgres.NewQueries(db)
	repo := repository.NewWorkersREPO(db, Queries)

	service := service.NewService(repo)

	r := RunService.NewRunService(service)

	log.Printf("Workers Service running on :%d port", cfg.ServicePort)
	if err := r.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
