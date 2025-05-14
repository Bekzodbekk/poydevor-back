package repository

import (
	"context"
	userpb "user-service/genproto/userspb"
)

type IUserRepository interface {
	Login(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginResp, error)
}
