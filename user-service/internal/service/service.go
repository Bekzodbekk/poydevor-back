package service

import (
	"context"
	"user-service/genproto/userspb"
	"user-service/internal/repository"
)

type Service struct {
	userspb.UnimplementedUsersServiceServer
	UserRepo repository.IUserRepository
}

func NewService(userRepo repository.IUserRepository) Service {
	return Service{
		UserRepo: userRepo,
	}
}
func (s *Service) Login(ctx context.Context, req *userspb.LoginReq) (*userspb.LoginResp, error) {
	return s.UserRepo.Login(ctx, req)
}
