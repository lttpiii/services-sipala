package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"services-sipala/services/users/types"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (c *controller) UpdateUser(ctx context.Context, req *types.ReqUpdateUser) (res *types.ResUpdateUser, err error) {
	if req.UserRole != "admin" && req.UserIDOnToken != req.UserID {
		return nil, fmt.Errorf("forbiden: admin required for update other user")
	}

	db := c.cfg.MysqlClient

	var (
		setClause []string
		args []any
	)

	if req.Name != "" {
		setClause = append(setClause, "name = ?")
		args = append(args, req.Name)
	}

	if req.Email != "" {
		setClause = append(setClause, "email = ?")
		args = append(args, req.Email)
	}

	if req.Role != "" && req.UserRole == "admin" {
		setClause = append(setClause, "role = ?")
		args = append(args, req.Role)
	}

	setClause = append(setClause, "updated_at = ?")
	args = append(args, time.Now())

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := fmt.Sprintf(`
		update users
		set %s
		where id = ? and deleted_at is null
	`, strings.Join(setClause, ", "))

	args = append(args, req.UserID)

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
	table := "users"
	desc := fmt.Sprintf("Activity recorded with action %s, change something on %s table", action, table)

	if _, err := tx.ExecContext(
		ctx,
		`
		insert into activity_logs (id, user_id, action, entity, entity_id, description, created_at)
		values (?,?,?,?,?,?,?)
		`,
		uuid.NewString(), req.UserIDOnToken, action, table, req.UserID, desc, time.Now(),
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}