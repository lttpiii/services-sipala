package controllers

import (
	"context"
	"fmt"
	"services-sipala/services/borrow/types"
	"time"

	"github.com/google/uuid"
)

func (c *controller) CreateBorrow(ctx context.Context, req *types.ReqCreateBorrow) (res *types.ResCreateBorrow, err error) {
	db := c.cfg.MysqlClient

	statusDefault := "pending"
	now := time.Now()

	if req.DueDate.Before(now) {
		return nil, fmt.Errorf("invalid request: duedate not valid")
	}

	id := uuid.NewString()

	if _, err := db.ExecContext(
		ctx,
		`
		insert into borrow_transactions (id, borrower_id, status, borrow_date, due_date, created_at)
		values (?,?,?,?,?,?)
		`, id, req.AuthUserID, statusDefault, now, req.DueDate, time.Now(),
	); err != nil {
		return nil, err
	}

	return &types.ResCreateBorrow{
		ID: id,
	}, nil
}