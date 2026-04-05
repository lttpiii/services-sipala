package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"services-sipala/services/borrow/types"
)

func (c *controller) RemoveBorrowItem(ctx context.Context, req *types.ReqRemoveBorrowItem) (res *types.ResRemoveBorrowItem, err error) {
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

	var ItemExists bool

	if err := tx.QueryRowContext(
		ctx,
		`
		select exists (
			select 1 from borrow_transaction_items
			where id = ?
		)
		`, req.ItemID,
	).Scan(&ItemExists); err != nil {
		return nil, err
	}

	if !ItemExists {
		return nil, fmt.Errorf("not found: data not found")
	}

	if req.AuthUserRole != "admin" && req.AuthUserRole != "staff" {
		var borrowerID string

		err := db.QueryRowContext(
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
		delete from borrow_transaction_items
		where id = ?
		`, req.ItemID,
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}