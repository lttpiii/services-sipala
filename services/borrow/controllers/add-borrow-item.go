package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"services-sipala/services/borrow/types"
	"time"

	"github.com/google/uuid"
)

func (c *controller) AddBorrowItem(ctx context.Context, req *types.ReqAddBorrowItem) (res *types.ResAddBorrowItem, err error) {
	db := c.cfg.MysqlClient

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var exists bool

	if err := tx.QueryRowContext(
		ctx,
		`
		select exists (
			select 1 from borrow_transactions
			where id = ? and status = 'pending'
		)
		`, req.BorrowID,
	).Scan(&exists); err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("not found: data not found")
	}

	var ToolExists bool

	if err := tx.QueryRowContext(
		ctx,
		`
		select exists (
			select 1 from tools
			where id = ? and deleted_at is null
		)
		`, req.ToolId,
	).Scan(&ToolExists); err != nil {
		return nil, err
	}

	if !ToolExists {
		return nil, fmt.Errorf("not found: data not found")
	}

	var availableStock int

	if err := tx.QueryRowContext(
		ctx,
		`
		select 
			t.stock - coalesce(sum(bti.quantity), 0) 
		from tools t
		left join borrow_transaction_items bti on bti.tool_id = t.id
		left join borrow_transactions bt 
			on bt.id = bti.borrow_transaction_id 
			and bt.status NOT IN ('returned', 'rejected')
		where t.id = ? and t.deleted_at is null
		group by t.id
		for update
		`, req.ToolId,
	).Scan(&availableStock); err != nil {
		return nil, err
	}

	if availableStock < req.Quantity {
		return nil, fmt.Errorf("invalid request: stock is not enough")
	}
	
	if req.AuthUserRole != "admin" && req.AuthUserRole != "staff" {
		var borrowerID string

		err := tx.QueryRowContext(
			ctx,
			`
			SELECT borrower_id
			FROM borrow_transactions
			WHERE id = ?
			`,
			req.BorrowID,
		).Scan(&borrowerID)

		if err != nil {
			return nil, err
		}

		if borrowerID != req.AuthUserID {
			return nil, fmt.Errorf("forbidden: user must be owner borrow")
		}
	}

	if _, err := tx.ExecContext(
		ctx,
		`
		insert into borrow_transaction_items (id, borrow_transaction_id, tool_id, quantity, created_at)
		values (?,?,?,?,?)
		`, uuid.NewString(), req.BorrowID, req.ToolId, req.Quantity, time.Now(),
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}