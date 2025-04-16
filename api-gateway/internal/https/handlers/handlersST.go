package handlers

import (
	"api-gateway/internal/pkg/service"
)

type HandlersST struct {
	service *service.ServiceRepositoryClient
}

func NewHandlers(service *service.ServiceRepositoryClient) *HandlersST {
	return &HandlersST{
		service: service,
	}
}
