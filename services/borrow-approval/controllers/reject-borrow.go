package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"services-sipala/services/borrow-approval/types"
	"time"

	"github.com/google/uuid"
)

func (c *controller) RejectBorrow(ctx context.Context, req *types.ReqRejectBorrow) (res *types.ResRejectBorrow, err error) {
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

	result, err := tx.ExecContext(
		ctx,
		`
		update borrow_transactions set
			status = 'rejected', approved_by = ?, approved_at = ?
		where id = ?
		`, req.AuthUserID, time.Now(), req.BorrowID,
	)

	rowsAffected, err := result.RowsAffected()
	if err != nil { 
		return nil, fmt.Errorf("not found: zero rows affected") 
	} 

	if rowsAffected == 0 { 
		return nil, fmt.Errorf("not found: zero rows affected") 
	}

	action := "REJECTED"
	table := "borrow_transactions"
	desc := fmt.Sprintf("Activity recorded with action %s, change something on %s table", action, table)

	if _, err := tx.ExecContext(
		ctx,
		`
		insert into activity_logs (id, user_id, action, entity, entity_id, description, created_at)
		values (?,?,?,?,?,?,?)
		`,
		uuid.NewString(), req.AuthUserID, action, table, req.BorrowID, desc, time.Now(),
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err 
	}

	return nil, nil
}