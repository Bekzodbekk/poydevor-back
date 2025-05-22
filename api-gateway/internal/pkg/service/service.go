package service

import (
	"api-gateway/genproto/userspb"
	"api-gateway/genproto/workerspb"
	"context"
)

type ServiceRepositoryClient struct {
	WorkersServiceClient workerspb.WorkersServiceClient
	UsersServiceClient   userspb.UsersServiceClient
}

func NewServiceRepositoryClient(
	workersServiceClient *workerspb.WorkersServiceClient,
	usersServiceClient *userspb.UsersServiceClient,
) *ServiceRepositoryClient {
	return &ServiceRepositoryClient{
		WorkersServiceClient: *workersServiceClient,
		UsersServiceClient:   *usersServiceClient,
	}
}

// Workers Service
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
func (s *ServiceRepositoryClient) UpdateWorker(ctx context.Context, req *workerspb.UpdateWorkerReq) (*workerspb.UpdateWorkerResp, error) {
	return s.WorkersServiceClient.UpdateWorker(ctx, req)
}
func (s *ServiceRepositoryClient) GetDailyProductionWorkersById(ctx context.Context, req *workerspb.GetDailyProductionWorkersByIdReq) (*workerspb.GetDailyProductionWorkersByIdResp, error) {
	return s.WorkersServiceClient.GetDailyProductionWorkersById(ctx, req)
}
func (s *ServiceRepositoryClient) GetLoadProductionWorkersById(ctx context.Context, req *workerspb.GetLoadProductionWorkersByIdReq) (*workerspb.GetLoadProductionWorkersByIdResp, error) {
	return s.WorkersServiceClient.GetLoadProductionWorkersById(ctx, req)
}

// Users Service
func (s *ServiceRepositoryClient) Login(ctx context.Context, req *userspb.LoginReq) (*userspb.LoginResp, error) {
	return s.UsersServiceClient.Login(ctx, req)
}
