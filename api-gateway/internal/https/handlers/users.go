package handlers

import (
	"api-gateway/genproto/userspb"
	authmiddleware "api-gateway/internal/https/middleware/AuthMiddleware"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *HandlersST) Login(ctx *gin.Context) {
	req := userspb.LoginReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	fmt.Println(req.Login)

	resp, err := h.service.UsersServiceClient.Login(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlersST) CheckToken(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	user, err := authmiddleware.ExtractClaim(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}
	ctx.JSON(200, gin.H{"user": user})
}
