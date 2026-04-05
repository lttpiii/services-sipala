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

func (c *controller) CreateUser(ctx context.Context, req *types.ReqCreateUser) (res *types.ResCreateUser, err error) {
	db := c.cfg.MysqlClient

	if strings.Contains(req.Password, " ") {
		return nil, fmt.Errorf("invalid request: invalid password format")
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	id := uuid.NewString()
	hash, err := c.utilities.HashString(req.Password)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	if _, err := tx.ExecContext(
		ctx,
		`
		insert into users (id, name, email, password_hash, role, created_at)
		values (?,?,?,?,?,?)
		`, id, req.Name, req.Email, hash, req.Role, now,
	); err != nil {
		return nil, err
	}

	action := "CREATE"
	table := "users"
	desc := fmt.Sprintf("Activity recorded with action %s, change something on %s table", action, table)

	if _, err := tx.ExecContext(
		ctx,
		`
		insert into activity_logs (id, user_id, action, entity, entity_id, description, created_at)
		values (?,?,?,?,?,?,?)
		`,
		uuid.NewString(), req.UserID, action, table, id, desc, time.Now(),
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}