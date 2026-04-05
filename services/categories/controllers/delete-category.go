package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"services-sipala/services/categories/types"
	"time"

	"github.com/google/uuid"
)

func (c *controller) DeleteCategory(ctx context.Context, req *types.ReqDeleteCategory) (res *types.ResDeleteCategory, err error) {
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
			select 1 from tools
			where category_id = ?
		)
		`, req.CategoryID,
	).Scan(&exists); err != nil {
		return nil, err
	}

	if exists {
		action := "DELETE"
		table := "categories"
		desc := fmt.Sprintf("Activity rejected with action %s", action)

		if _, err := tx.ExecContext(
			ctx,
			`
			insert into activity_logs (id, user_id, action, entity, entity_id, description, created_at)
			values (?,?,?,?,?,?,?)
			`,
			uuid.NewString(), req.AuthUserID, action, table, req.CategoryID, desc, time.Now(),
		); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("conflict: category still has relationships")
	}

	result, err := tx.ExecContext(
		ctx,
		`
		delete from categories
		where id = ?
		`, req.CategoryID,
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
	table := "categories"
	desc := fmt.Sprintf("Activity recorded with action %s, change something on %s table", action, table)

	if _, err := tx.ExecContext(
		ctx,
		`
		insert into activity_logs (id, user_id, action, entity, entity_id, description, created_at)
		values (?,?,?,?,?,?,?)
		`,
		uuid.NewString(), req.AuthUserID, action, table, req.CategoryID, desc, time.Now(),
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}