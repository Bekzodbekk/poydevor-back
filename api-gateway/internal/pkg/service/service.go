package service

import (
	"api-gateway/genproto/workerspb"
	"context"
)

type ServiceRepositoryClient struct {
	WorkersServiceClient workerspb.WorkersServiceClient
}

func NewServiceRepositoryClient(workersServiceClient *workerspb.WorkersServiceClient) *ServiceRepositoryClient {
	return &ServiceRepositoryClient{
		WorkersServiceClient: *workersServiceClient,
	}
}

func (s *ServiceRepositoryClient) AddWorkers(ctx context.Context, req *workerspb.AddWorkersReq) (*workerspb.AddWorkersResp, error) {
	return s.WorkersServiceClient.AddWorkers(ctx, req)
}

func (s *ServiceRepositoryClient) GetWorkers(ctx context.Context, req *workerspb.GetWorkersReq) (*workerspb.GetWorkersResp, error) {
	return s.WorkersServiceClient.GetWorkers(ctx, req)
}
func (s *ServiceRepositoryClient) EndDay(ctx context.Context, req *workerspb.EndDayReq) (*workerspb.EndDayResp, error) {
	return s.WorkersServiceClient.EndDay(ctx, req)
}
func (s *ServiceRepositoryClient) LoadBlocks(ctx context.Context, req *workerspb.LoadBlocksReq) (*workerspb.LoadBlocksResp, error) {
	return s.WorkersServiceClient.LoadBlocks(ctx, req)
}
func (s *ServiceRepositoryClient) MonthlyReport(ctx context.Context, req *workerspb.MonthlyReportReq) (*workerspb.MonthlyReportResp, error) {
	return s.WorkersServiceClient.MonthlyReport(ctx, req)
}
func (s *ServiceRepositoryClient) AddPaidMonthly(ctx context.Context, req *workerspb.PaidWorkerMonthlyReq) (*workerspb.PaidWorkerMonthlyResp, error) {
	return s.WorkersServiceClient.AddPaidMonthly(ctx, req)
}
