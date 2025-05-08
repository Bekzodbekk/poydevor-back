package repository

import (
	"context"
	userpb "user-service/genproto/userspb"
)

type IUserRepository interface {
	Login(cfg context.Context, req *userpb.LoginReq) (*userpb.LoginResp, error)
}
