package handlers

import (
	"api-gateway/genproto/userspb"

	"github.com/gin-gonic/gin"
)

func (h *HandlersST) Login(ctx *gin.Context) {
	req := userspb.LoginReq{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.service.UsersServiceClient.Login(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
