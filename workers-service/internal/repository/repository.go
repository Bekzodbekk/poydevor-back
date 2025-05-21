package repository

import (
	"context"
	"workers-service/genproto/workerspb"
)

type IWorkersRepository interface {
	AddWorkers(ctx context.Context, req *workerspb.AddWorkersReq) (*workerspb.AddWorkersResp, error)
	GetWorkers(ctx context.Context, req *workerspb.GetWorkersReq) (*workerspb.GetWorkersResp, error)
	EndDay(ctx context.Context, req *workerspb.EndDayReq) (*workerspb.EndDayResp, error)
	LoadBlocks(ctx context.Context, req *workerspb.LoadBlocksReq) (*workerspb.LoadBlocksResp, error)
	MonthlyReport(ctx context.Context, req *workerspb.MonthlyReportReq) (*workerspb.MonthlyReportResp, error)
	AddPaidMonthly(ctx context.Context, req *workerspb.PaidWorkerMonthlyReq) (*workerspb.PaidWorkerMonthlyResp, error)
	UpdateWorker(ctx context.Context, req *workerspb.UpdateWorkerReq) (*workerspb.UpdateWorkerResp, error)
	DeleteWorker(ctx context.Context, req *workerspb.DeleteWorkerReq) (*workerspb.DeleteWorkerResp, error)
	GetDailyProductionWorkersById(ctx context.Context, req *workerspb.GetDailyProductionWorkersByIdReq) (*workerspb.GetDailyProductionWorkersByIdResp, error)
}
