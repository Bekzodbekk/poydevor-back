package repository

import (
	"context"
	"database/sql"
	"strconv"
	"time"
	"workers-service/genproto/workerspb"
	"workers-service/storage"
)

type WorkersREPO struct {
	basicDB *sql.DB
	queries *storage.Queries
}

func NewWorkersREPO(db *sql.DB, queries *storage.Queries) *WorkersREPO {
	return &WorkersREPO{
		basicDB: db,
		queries: queries,
	}
}

func (w *WorkersREPO) AddWorkers(ctx context.Context, req *workerspb.AddWorkersReq) (*workerspb.AddWorkersResp, error) {
	err := w.queries.AddWorker(ctx, storage.AddWorkerParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}

	return &workerspb.AddWorkersResp{
		Status:  true,
		Message: "Worker added successfully",
	}, nil
}
func (w *WorkersREPO) GetWorkers(ctx context.Context, req *workerspb.GetWorkersReq) (*workerspb.GetWorkersResp, error) {
	resp, err := w.queries.GetWorkers(ctx)
	if err != nil {
		return nil, err
	}

	workers := []*workerspb.Worker{}

	for _, elm := range resp {
		id := strconv.Itoa(int(elm.ID))
		worker := workerspb.Worker{
			Id:          id,
			FirstName:   elm.FirstName,
			LastName:    elm.LastName,
			PhoneNumber: elm.Phone,
		}
		workers = append(workers, &worker)
	}

	return &workerspb.GetWorkersResp{
		Workers: workers,
	}, nil
}

func (w *WorkersREPO) EndDay(ctx context.Context, req *workerspb.EndDayReq) (*workerspb.EndDayResp, error) {
	tx, err := w.basicDB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		tx.Rollback()
	}()

	sqlcTX := w.queries.WithTx(tx)

	t, err := time.Parse("2006-01-02", req.Date)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := sqlcTX.EndDay(ctx, storage.EndDayParams{
		Date:        t,
		CountBlocks: req.CountBlocks,
	})

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, elm := range req.Workers {

		err := w.queries.EndDayWorkers(ctx, storage.EndDayWorkersParams{
			DailyProductionID: id,
			WorkerID:          elm,
		})
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &workerspb.EndDayResp{
		Status:  true,
		Message: "End day successfully",
	}, nil
}

func (w *WorkersREPO) LoadBlocks(ctx context.Context, req *workerspb.LoadBlocksReq) (*workerspb.LoadBlocksResp, error) {
	tx, err := w.basicDB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		tx.Rollback()
	}()
	sqlcTX := w.queries.WithTx(tx)
	t, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := sqlcTX.SendBlocks(ctx, storage.SendBlocksParams{
		Date:        t,
		CountBlocks: int32(req.CountBlocks),
		Address:     req.Address,
		LoadPrice:   int32(req.PriceBlock),
	})

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, elm := range req.Workers {
		err := sqlcTX.LoadBlockWorkers(ctx, storage.LoadBlockWorkersParams{
			SendBlockID: id,
			WorkerID:    elm,
		})
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &workerspb.LoadBlocksResp{
		Status:  true,
		Message: "Blocks loaded successfully",
	}, nil
}

func (w *WorkersREPO) MonthlyReport(ctx context.Context, req *workerspb.MonthlyReportReq) (*workerspb.MonthlyReportResp, error) {
	worker_id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	year, err := strconv.Atoi(req.Year)
	if err != nil {
		return nil, err
	}

	month, err := strconv.Atoi(req.Month)
	if err != nil {
		return nil, err
	}
	endDayResp, err := w.queries.EndDayDataMonthlyReport(ctx, storage.EndDayDataMonthlyReportParams{
		Column1: int32(worker_id),
		Column2: int32(year),
		Column3: int32(month),
	})

	if err != nil {
		return nil, err
	}
	endDayData := []*workerspb.EndDayData{}

	for _, elm := range endDayResp {
		id := strconv.Itoa(int(elm.DailyProductionID))
		endDay := workerspb.EndDayData{
			Id:           id,
			Date:         elm.Date.GoString(),
			CountBlocks:  elm.TotalBlocks,
			WorkersCount: int32(elm.WorkerCount),
			WorkerShare:  elm.WorkerShare,
			TotalPrice:   elm.WorkerPayment,
		}
		endDayData = append(endDayData, &endDay)
	}

	loadBlocksData := []*workerspb.LoadBlocksData{}
	loadBlocksDataResp, err := w.queries.LoadBlocksDataMonthlyReport(ctx, storage.LoadBlocksDataMonthlyReportParams{
		Column1: int32(worker_id),
		Column2: int32(year),
		Column3: int32(month),
	})
	if err != nil {
		return nil, err
	}
	for _, elm := range loadBlocksDataResp {
		id := strconv.Itoa(int(elm.SendBlockID))
		loadBlocks := workerspb.LoadBlocksData{
			Id:              id,
			Date:            elm.Date.GoString(),
			CountBlocks:     elm.TotalBlocks,
			WorkersCount:    int32(elm.WorkerCount),
			BlocksPerWorker: elm.BlocksPerWorker,
			TotalPrice:      elm.Payment,
		}
		loadBlocksData = append(loadBlocksData, &loadBlocks)
	}

	paidMonthlyData := []*workerspb.PaidMonthly{}
	paidMonthlyDataResp, err := w.queries.PaidMonthlyData(ctx, storage.PaidMonthlyDataParams{
		Column1: int32(worker_id),
		Column2: int32(year),
		Column3: int32(month),
	})
	if err != nil {
		return nil, err
	}
	for _, elm := range paidMonthlyDataResp {
		paidMonthly := workerspb.PaidMonthly{
			Date:      elm.Date.GoString(),
			PaidPrice: int64(elm.PaidPrice),
		}
		paidMonthlyData = append(paidMonthlyData, &paidMonthly)
	}
	return &workerspb.MonthlyReportResp{
		EndDayData:     endDayData,
		LoadBlocksData: loadBlocksData,
		PaidMonthly:    paidMonthlyData,
	}, nil
}

func (w *WorkersREPO) AddPaidMonthly(ctx context.Context, req *workerspb.PaidWorkerMonthlyReq) (*workerspb.PaidWorkerMonthlyResp, error) {
	err := w.queries.AddPaidMonthly(ctx, storage.AddPaidMonthlyParams{
		WorkerID:  int32(req.WorkerId),
		Date:      time.Now(),
		PaidPrice: float64(req.PaidPrice),
	})
	if err != nil {
		return nil, err
	}

	return &workerspb.PaidWorkerMonthlyResp{
		Status:  true,
		Message: "Paid successfully",
	}, nil
}

func (w *WorkersREPO) UpdateWorker(ctx context.Context, req *workerspb.UpdateWorkerReq) (*workerspb.UpdateWorkerResp, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = w.queries.UpdateWorkers(ctx, storage.UpdateWorkersParams{
		ID:        int32(id),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
	})

	if err != nil {
		return nil, err
	}

	return &workerspb.UpdateWorkerResp{
		Status:  true,
		Message: "Update worker Successfuly",
	}, nil

}

func (w *WorkersREPO) DeleteWorker(ctx context.Context, req *workerspb.DeleteWorkerReq) (*workerspb.DeleteWorkerResp, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = w.queries.DeleteWorkers(ctx, storage.DeleteWorkersParams{
		ID:        int32(id),
		DeletedAt: sql.NullInt32{Int32: int32(time.Now().Unix()), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &workerspb.DeleteWorkerResp{
		Status:  true,
		Message: "Delete Worker Successfuly",
	}, nil

}

func (w *WorkersREPO) GetDailyProductionWorkersById(ctx context.Context, req *workerspb.GetDailyProductionWorkersByIdReq) (*workerspb.GetDailyProductionWorkersByIdResp, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := w.queries.GetDailyProductionWorkersNameById(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	workersData := []*workerspb.Worker{}
	for _, elm := range resp {
		id := strconv.Itoa(int(elm.ID))
		worker := workerspb.Worker{
			Id:        id,
			FirstName: elm.FirstName,
			LastName:  elm.LastName,
		}

		workersData = append(workersData, &worker)
	}

	return &workerspb.GetDailyProductionWorkersByIdResp{
		Status:  true,
		Message: "Workers GET Successfuly",
		Worker:  workersData,
	}, nil
}
func (w *WorkersREPO) GetLoadProductionWorkersById(ctx context.Context, req *workerspb.GetLoadProductionWorkersByIdReq) (*workerspb.GetLoadProductionWorkersByIdResp, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := w.queries.GetLoadProductionWorkersNameById(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	workersData := []*workerspb.Worker{}
	for _, elm := range resp {
		id := strconv.Itoa(int(elm.ID))
		worker := workerspb.Worker{
			Id:        id,
			FirstName: elm.FirstName,
			LastName:  elm.LastName,
		}

		workersData = append(workersData, &worker)
	}

	return &workerspb.GetLoadProductionWorkersByIdResp{
		Status:  true,
		Message: "Workers GET Successfuly",
		Worker:  workersData,
	}, nil
}
