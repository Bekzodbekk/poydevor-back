package main

import (
	"api-gateway/internal/https"
	"api-gateway/internal/pkg/config"
	"api-gateway/internal/pkg/service"
	workersservice "api-gateway/internal/pkg/workers-service"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {

		if err := r.ListenAndServeTLS(cfg.TLS.CertFile, cfg.TLS.KeyFile); err != nil {
			log.Fatal(err)
		}
	}()
	signalReceived := <-sigChan
	log.Print("Received signal", signalReceived)
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := r.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server shutdown error: ", err)
	}
	log.Print("Graceful shutdown complete.")
}
