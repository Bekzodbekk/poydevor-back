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
	return nil, nil
}
