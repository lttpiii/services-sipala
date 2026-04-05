package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"services-sipala/services/returns/types"
	"time"
)

func (c *controller) CalculateFine(ctx context.Context, req *types.ReqCalculateFineReturns) (res *types.ResCalculateFineReturns, err error) {
	db := c.cfg.MysqlClient

	now := time.Now()

	var (
		borrowID string
		dueDate  time.Time
		status   string
	)

	err = db.QueryRowContext(ctx, `
		SELECT 
			id,
			due_date,
			status
		FROM borrow_transactions
		WHERE id = ?
	`, req.BorrowTransactionID).Scan(
		&borrowID,
		&dueDate,
		&status,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("borrow transaction not found")
	}
	if err != nil {
		return nil, err
	}

	if status != "approved" {
		return nil, fmt.Errorf("borrow status must be approved")
	}

	lateDays := int(math.Max(0, now.Sub(dueDate).Hours()/24))

	const finePerDay = 10000
	fineAmount := float64(lateDays * finePerDay)

	res = &types.ResCalculateFineReturns{
		BorrowTransactionID: borrowID,
		DueDate:             dueDate,
		CurrentDate:         now,
		LateDays:            lateDays,
		FinePerDay:          finePerDay,
		FineAmount:          fineAmount,
		Currency:            "IDR",
	}

	return res, nil
}