package handlers

import (
	"api-gateway/genproto/workerspb"
	"strconv"

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

func (h *HandlersST) MonthlyReport(ctx *gin.Context) {
	workerId := ctx.Param("worker_id")

	resp, err := h.service.MonthlyReport(ctx, &workerspb.MonthlyReportReq{
		Id: workerId,
	})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlersST) AddPaidMonthly(ctx *gin.Context) {
	workerId := ctx.Param("worker_id")
	req := workerspb.PaidWorkerMonthlyReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	worker_id, err := strconv.Atoi(workerId)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	req.WorkerId = int64(worker_id)

	resp, err := h.service.AddPaidMonthly(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
