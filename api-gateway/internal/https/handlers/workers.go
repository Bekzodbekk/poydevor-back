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
	workerId := ctx.Query("worker_id")
	month := ctx.Query("month")
	year := ctx.Query("year")

	resp, err := h.service.MonthlyReport(ctx, &workerspb.MonthlyReportReq{
		Id:    workerId,
		Month: month,
		Year:  year,
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

func (h *HandlersST) UpdateWorker(ctx *gin.Context) {
	id := ctx.Param("worker_id")
	worker := workerspb.UpdateWorkerReq{}
	if err := ctx.BindJSON(&worker); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	worker.Id = id
	resp, err := h.service.UpdateWorker(ctx, &worker)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlersST) DeleteWorker(ctx *gin.Context) {
	id := ctx.Param("worker_id")
	resp, err := h.service.WorkersServiceClient.DeleteWorker(ctx, &workerspb.DeleteWorkerReq{
		Id: id,
	})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, resp)
}
