package workersservice

import (
	"api-gateway/genproto/userspb"
	"api-gateway/internal/pkg/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithUsersService(cfg config.Config) (*userspb.UsersServiceClient, error) {
	target := fmt.Sprintf("%s:%d", cfg.Services.UserService.Host, cfg.Services.UserService.Port)
	fmt.Println(target)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	usersServiceClient := userspb.NewUsersServiceClient(conn)
	return &usersServiceClient, nil
}
