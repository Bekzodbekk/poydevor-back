package service

import (
	"fmt"
	"net"
	"workers-service/genproto/workerspb"
	"workers-service/internal/pkg/config"
	"workers-service/internal/service"

	"google.golang.org/grpc"
)

type RunService struct {
	srv *service.Service
}

func NewRunService(srv *service.Service) *RunService {
	return &RunService{
		srv: srv,
	}
}

func (r *RunService) Run(cfg *config.Config) error {
	target := fmt.Sprintf("%s:%d", cfg.ServiceHost, cfg.ServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	newServe := grpc.NewServer()
	workerspb.RegisterWorkersServiceServer(newServe, r.srv)
	return newServe.Serve(listener)
}
