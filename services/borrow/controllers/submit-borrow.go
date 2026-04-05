package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"services-sipala/services/borrow/types"
)

func (c *controller) SubmitBorrow(ctx context.Context, req *types.ReqSubmitBorrow) (res *types.ResSubmitBorrow, err error) {
	db := c.cfg.MysqlClient

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

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

	var exists bool

	if err := tx.QueryRowContext(
		ctx,
		`
		select exists (
			select 1 from borrow_transactions
			where id = ?
		)
		`, req.BorrowID,
	).Scan(&exists); err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("not found: data not found")
	}

	var itemsExists bool

	if err := tx.QueryRowContext(
		ctx,
		`
		select exists (
			select 1 from borrow_transaction_items
			where borrow_transaction_id = ?
		)	
		`, req.BorrowID,
	).Scan(&itemsExists); err != nil {
		return nil, err
	}

	if !itemsExists {
		return nil, fmt.Errorf("not found: please add item first")
	}

	_, err = tx.ExecContext(
		ctx,
		`
		update borrow_transactions
		set status = 'pending'
		where id = ?
		`, req.BorrowID,
	)

	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}