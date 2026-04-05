package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"services-sipala/services/tools/types"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (c *controller) UpdateTool(ctx context.Context, req *types.ReqUpdateTool) (res *types.ResUpdateTool, err error) {
	db := c.cfg.MysqlClient

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var (
		setClause []string
		args []any
	)

	if req.Name != "" {
		setClause = append(setClause, "name = ?")
		args = append(args, req.Name)
	}

	if req.CategoryID != "" {
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

		setClause = append(setClause, "category_id = ?")
		args = append(args, req.CategoryID)
	}

	if req.Stock >= 0 {
		setClause = append(setClause, "stock = ?")
		args = append(args, req.Stock)
	}

	if req.Description != "" {
		setClause = append(setClause, "description = ?")
		args = append(args, req.Description)
	}

	setClause = append(setClause, "updated_at = ?")
	args = append(args, time.Now())

	query := fmt.Sprintf(`
		update tools
		set %s
		where id = ? and deleted_at is null
	`, strings.Join(setClause, ", "))

	args = append(args, req.ToolID)

	result, err := tx.ExecContext(
		ctx,
		query,
		args...
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
	table := "tools"
	desc := fmt.Sprintf("Activity recorded with action %s, change something on %s table", action, table)

	if _, err := tx.ExecContext(
		ctx,
		`
		insert into activity_logs (id, user_id, action, entity, entity_id, description, created_at)
		values (?,?,?,?,?,?,?)
		`,
		uuid.NewString(), req.AuthUserID, action, table, req.ToolID, desc, time.Now(),
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}