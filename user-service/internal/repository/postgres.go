package repository

import (
	"context"
	"database/sql"
	"errors"
	"user-service/genproto/userspb"
	"user-service/storage"
	"user-service/token"
)

type UserREPO struct {
	basicDB *sql.DB
	queries *storage.Queries
}

func NewUserREPO(db *sql.DB, queries *storage.Queries) *UserREPO {
	return &UserREPO{
		basicDB: db,
		queries: queries,
	}
}

func (u *UserREPO) Login(ctx context.Context, req *userspb.LoginReq) (*userspb.LoginResp, error) {
	resp, err := u.queries.Login(ctx, req.Login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &userspb.LoginResp{
				Status:  false,
				Message: "Login xato foydalanuvchi topilmadi!",
			}, nil
		}
		return nil, err
	}
	if resp.Password != req.Password {
		return &userspb.LoginResp{
			Status:  false,
			Message: "Parol xato!",
		}, nil
	}

	accessToken, err := token.CreateJWTToken(req.Id, "PoydevorAdmin")
	if err != nil {
		return nil, err
	}

	return &userspb.LoginResp{
		Status:  true,
		Message: "Login muvaffaqiyatli",
		Token:   accessToken,
	}, nil
}
