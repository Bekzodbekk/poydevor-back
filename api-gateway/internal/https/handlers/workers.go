package handlers

import (
	"api-gateway/genproto/workerspb"

	"github.com/gin-gonic/gin"
)

func (h *HandlersST) AddWorkers(ctx *gin.Context) {
	req := &workerspb.AddWorkersReq{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	resp, err := h.service.AddWorkers(ctx, req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlersST) AllWorkers(ctx *gin.Context) {
	req := &workerspb.GetWorkersReq{}

	resp, err := h.service.GetWorkers(ctx, req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlersST) EndDay(ctx *gin.Context) {
	req := workerspb.EndDayReq{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.service.EndDay(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlersST) LoadBlocks(ctx *gin.Context) {
	req := workerspb.LoadBlocksReq{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.service.LoadBlocks(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
