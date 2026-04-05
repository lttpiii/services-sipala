package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"services-sipala/services/users/types"
	"time"

	"github.com/google/uuid"
)

func (c *controller) DeleteUser(ctx context.Context, req *types.ReqDeleteUser) (res *types.ResDeleteUser, err error) {
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
			where borrower_id = ? and status <> 'returned'
		)
		`, req.UserID,
	).Scan(&exists); err != nil {
		return nil, err
	}

	if exists {
		action := "DELETE"
		table := "users"
		desc := fmt.Sprintf("Activity rejected with action %s", action)

		if _, err := tx.ExecContext(
			ctx,
			`
			insert into activity_logs (id, user_id, action, entity, entity_id, description, created_at)
			values (?,?,?,?,?,?,?)
			`,
			uuid.NewString(), req.AuthUserID, action, table, req.UserID, desc, time.Now(),
		); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("conflict: loan not yet returned")
	}

	result, err := tx.ExecContext(
		ctx,
		`
		update users
		set deleted_at = ?
		where id = ? and deleted_at is null	
		`, time.Now(), req.UserID,
	)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("not found: zero rows affected")
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("not found: zero rows affected")
	}

	action := "DELETE"
	table := "users"
	desc := fmt.Sprintf("Activity recorded with action %s, change something on %s table", action, table)

	if _, err := tx.ExecContext(
		ctx,
		`
		insert into activity_logs (id, user_id, action, entity, entity_id, description, created_at)
		values (?,?,?,?,?,?,?)
		`,
		uuid.NewString(), req.AuthUserID, action, table, req.UserID, desc, time.Now(),
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}