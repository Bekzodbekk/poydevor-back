package workersservice

import (
	"api-gateway/genproto/workerspb"
	"api-gateway/internal/pkg/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithWorkersService(cfg config.Config) (*workerspb.WorkersServiceClient, error) {
	target := fmt.Sprintf("%s:%d", cfg.Services.WorkersService.Host, cfg.Services.WorkersService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	workersServiceClient := workerspb.NewWorkersServiceClient(conn)
	return &workersServiceClient, nil
}
