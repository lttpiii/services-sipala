package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"services-sipala/services/returns/types"
)

func (c *controller) GetReturnByID(ctx context.Context, req *types.ReqGetReturnByID) (res *types.ResGetReturnByID, err error) {
	db := c.cfg.MysqlClient

	// =============================
	// 1. Ambil header return
	// =============================
	queryHeader := `
		SELECT
			rt.id,
			rt.borrow_transaction_id,
			rt.returned_at,
			rt.late_days,
			rt.fine_amount,
			rt.created_at,

			u.id,
			u.name,

			bu.name,
			bt.borrow_date,
			bt.due_date

		FROM return_transactions rt
		JOIN users u ON u.id = rt.processed_by
		JOIN borrow_transactions bt ON bt.id = rt.borrow_transaction_id
		JOIN users bu ON bu.id = bt.borrower_id

		WHERE rt.id = ?
	`

	res = &types.ResGetReturnByID{}

	var lateDays sql.NullInt64
	var fineAmount sql.NullFloat64

	err = db.QueryRowContext(ctx, queryHeader, req.ReturnID).Scan(
		&res.ID,
		&res.BorrowTransactionID,
		&res.ReturnedAt,
		&lateDays,
		&fineAmount,
		&res.CreatedAt,

		&res.ProcessedBy.ID,
		&res.ProcessedBy.Name,

		&res.BorrowDetails.BorrowerName,
		&res.BorrowDetails.BorrowDate,
		&res.BorrowDetails.DueDate,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("return not found")
		}
		return nil, err
	}

	// handle nullable
	if lateDays.Valid {
		val := int(lateDays.Int64)
		res.LateDays = &val
	}
	if fineAmount.Valid {
		val := float64(fineAmount.Float64)
		res.FineAmount = &val
	}

	// =============================
	// 2. Ambil items
	// =============================
	queryItems := `
		SELECT
			t.name,
			bti.quantity
		FROM borrow_transaction_items bti
		JOIN tools t ON t.id = bti.tool_id
		WHERE bti.borrow_transaction_id = ?
	`

	rows, err := db.QueryContext(ctx, queryItems, res.BorrowTransactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []types.BorrowItem

	for rows.Next() {
		var item types.BorrowItem

		err := rows.Scan(
			&item.ToolName,
			&item.Quantity,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	res.BorrowDetails.Items = items

	return res, nil
}