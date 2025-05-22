package service

import (
	"context"
	"workers-service/genproto/workerspb"
	"workers-service/internal/repository"
)

type Service struct {
	workerspb.UnimplementedWorkersServiceServer
	repo repository.IWorkersRepository
}

func NewService(repo repository.IWorkersRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddWorkers(ctx context.Context, req *workerspb.AddWorkersReq) (*workerspb.AddWorkersResp, error) {
	return s.repo.AddWorkers(ctx, req)
}

func (s *Service) GetWorkers(ctx context.Context, req *workerspb.GetWorkersReq) (*workerspb.GetWorkersResp, error) {
	return s.repo.GetWorkers(ctx, req)
}

func (s *Service) EndDay(ctx context.Context, req *workerspb.EndDayReq) (*workerspb.EndDayResp, error) {
	return s.repo.EndDay(ctx, req)
}

func (s *Service) LoadBlocks(ctx context.Context, req *workerspb.LoadBlocksReq) (*workerspb.LoadBlocksResp, error) {
	return s.repo.LoadBlocks(ctx, req)
}

func (s *Service) MonthlyReport(ctx context.Context, req *workerspb.MonthlyReportReq) (*workerspb.MonthlyReportResp, error) {
	return s.repo.MonthlyReport(ctx, req)
}

func (s *Service) AddPaidMonthly(ctx context.Context, req *workerspb.PaidWorkerMonthlyReq) (*workerspb.PaidWorkerMonthlyResp, error) {
	return s.repo.AddPaidMonthly(ctx, req)
}
func (s *Service) UpdateWorker(ctx context.Context, req *workerspb.UpdateWorkerReq) (*workerspb.UpdateWorkerResp, error) {
	return s.repo.UpdateWorker(ctx, req)
}
func (s *Service) DeleteWorker(ctx context.Context, req *workerspb.DeleteWorkerReq) (*workerspb.DeleteWorkerResp, error) {
	return s.repo.DeleteWorker(ctx, req)
}
func (s *Service) GetDailyProductionWorkersById(ctx context.Context, req *workerspb.GetDailyProductionWorkersByIdReq) (*workerspb.GetDailyProductionWorkersByIdResp, error) {
	return s.repo.GetDailyProductionWorkersById(ctx, req)
}
func (s *Service) GetLoadProductionWorkersById(ctx context.Context, req *workerspb.GetLoadProductionWorkersByIdReq) (*workerspb.GetLoadProductionWorkersByIdResp, error) {
	return s.repo.GetLoadProductionWorkersById(ctx, req)
}
