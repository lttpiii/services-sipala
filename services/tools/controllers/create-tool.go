package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"services-sipala/services/tools/types"

	"time"

	"github.com/google/uuid"
)

func (c *controller) CreateTool(ctx context.Context, req *types.ReqCreateTool) (res *types.ResCreateTool, err error) {
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
			select 1 from categories
			where id = ?
		)
		`, req.CategoryID,
	).Scan(&exists); err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("not found: category not found")
	}

	id := uuid.NewString()
	now := time.Now()

	if _, err := tx.ExecContext(
		ctx,
		`
		insert into tools (id, name, category_id, stock, description, created_at)
		values (?,?,?,?,?,?)
		`, id, req.Name, req.CategoryID, req.Stock, req.Description, now,
	); err != nil {
		return nil, err
	}

	action := "CREATE"
	table := "tools"
	desc := fmt.Sprintf("Activity recorded with action %s, change something on %s table", action, table)

	if _, err := tx.ExecContext(
		ctx,
		`
		insert into activity_logs (id, user_id, action, entity, entity_id, description, created_at)
		values (?,?,?,?,?,?,?)
		`,
		uuid.NewString(), req.AuthUserID, action, table, id, desc, time.Now(),
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}