package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"services-sipala/services/returns/types"
	"time"

	"github.com/google/uuid"
)

func (c *controller) CreateReturns(ctx context.Context, req *types.ReqCreateReturns) (res *types.ResCreateReturns, err error) {
	db := c.cfg.MysqlClient

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	now := time.Now()

	// validasi borrow exists
	var (
		borrowID     string
		borrowerName string
		borrowDate   time.Time
		dueDate      time.Time
		status       string
	)

	err = tx.QueryRowContext(ctx, `
		SELECT 
			bt.id,
			u.name,
			bt.borrow_date,
			bt.due_date,
			bt.status
		FROM borrow_transactions bt
		JOIN users u ON u.id = bt.user_id
		WHERE bt.id = ?
		FOR UPDATE
	`, req.BorrowTransactionID).Scan(
		&borrowID,
		&borrowerName,
		&borrowDate,
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

	// cek apakah sudah ada return
	var exists bool
	err = tx.QueryRowContext(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM return_transactions WHERE borrow_transaction_id = ?
		)
	`, req.BorrowTransactionID).Scan(&exists)

	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("return transaction already exists")
	}

	// hitung late fine
	lateDays := int(math.Max(0, now.Sub(dueDate).Hours()/24))

	const fineRate = 10000 // per hari
	fineAmount := float64(lateDays * fineRate)

	// insert return
	returnID := uuid.New().String()

	_, err = tx.ExecContext(ctx, `
		INSERT INTO return_transactions (
			id,
			borrow_transaction_id,
			returned_at,
			late_days,
			fine_amount,
			processed_by,
			created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`,
		returnID,
		req.BorrowTransactionID,
		now,
		lateDays,
		fineAmount,
		req.AuthUserID,
		now,
	)

	if err != nil {
		return nil, err
	}

	// update status borrow
	_, err = tx.ExecContext(ctx, `
		UPDATE borrow_transactions
		SET status = 'returned'
		WHERE id = ?
	`, req.BorrowTransactionID)

	if err != nil {
		return nil, err
	}

	// get items
	rows, err := tx.QueryContext(ctx, `
		SELECT 
			t.id,
			t.name,
			bti.quantity
		FROM borrow_transaction_items bti
		JOIN tools t ON t.id = bti.tool_id
		WHERE bti.borrow_transaction_id = ?
	`, req.BorrowTransactionID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []types.BorrowItem

	for rows.Next() {
		var (
			toolID   string
			toolName string
			qty      int
		)

		if err := rows.Scan(&toolID, &toolName, &qty); err != nil {
			return nil, err
		}

		items = append(items, types.BorrowItem{
			ToolName: toolName,
			Quantity: qty,
		})
	}

	// get processed by
	var processedBy types.ProcessedBy
	err = tx.QueryRowContext(ctx, `
		SELECT id, name
		FROM users
		WHERE id = ?
	`, req.AuthUserID).Scan(&processedBy.ID, &processedBy.Name)

	if err != nil {
		return nil, err
	}

	action := "CREATE"
	table := "return_transactions"
	desc := fmt.Sprintf("Activity recorded with action %s, change something on %s table", action, table)

	if _, err := tx.ExecContext(
		ctx,
		`
		insert into activity_logs (id, user_id, action, entity, entity_id, description, created_at)
		values (?,?,?,?,?,?,?)
		`,
		uuid.NewString(), req.AuthUserID, action, table, returnID, desc, time.Now(),
	); err != nil {
		return nil, err
	}

	// commit
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// build response
	res = &types.ResCreateReturns{
		ID:                  returnID,
		BorrowTransactionID: req.BorrowTransactionID,
		ReturnedAt:          now,
		LateDays:            lateDays,
		FineAmount:          fineAmount,
		ProcessedBy:         processedBy,
		BorrowDetails: types.BorrowDetails{
			BorrowerName: borrowerName,
			BorrowDate:   borrowDate,
			DueDate:      dueDate,
			Items:        items,
		},
		CreatedAt: now,
	}

	return res, nil
}