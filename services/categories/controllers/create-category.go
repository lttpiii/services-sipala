package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"services-sipala/services/categories/types"
	"time"

	"github.com/google/uuid"
)

func (c *controller) CreateCategory(ctx context.Context, req *types.ReqCreateCategory) (res *types.ResCreateCategory, err error) {
	db := c.cfg.MysqlClient

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	id := uuid.NewString()
	now := time.Now()

	if _, err := tx.ExecContext(
		ctx,
		`
		insert into categories (id, name, created_at)
		values (?,?,?)
		`, id, req.Name, now,
	); err != nil {
		return nil, err
	}

	action := "CREATE"
	table := "categories"
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