package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"services-sipala/services/categories/types"
	"time"

	"github.com/google/uuid"
)

func (c *controller) UpdateCategory(ctx context.Context, req *types.ReqUpdateCategory) (res *types.ResUpdateCategory, err error) {
	db := c.cfg.MysqlClient

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(
		ctx,
		`
		update categories
		set name = ?
		where id = ?
		`, req.Name, req.CategoryID,
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

	action := "UPDATE"
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