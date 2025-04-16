package main

import (
	"api-gateway/internal/https"
	"api-gateway/internal/pkg/config"
	"api-gateway/internal/pkg/service"
	workersservice "api-gateway/internal/pkg/workers-service"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LOAD("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	conn1, err := workersservice.DialWithWorkersService(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	service := service.NewServiceRepositoryClient(conn1)

	r := https.NewGin(service)
	target := fmt.Sprintf("%s:%d", cfg.Services.ApiGateway.Host, cfg.Services.ApiGateway.Port)
	fmt.Println(target)
	log.Printf("Api Gateway running on :%d port", cfg.Services.ApiGateway.Port)
	if err := r.Run(target); err != nil {
		log.Fatal(err)
	}
}
