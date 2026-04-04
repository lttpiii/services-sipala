package controllers

import (
	"context"
	"fmt"
	"services-sipala/services/authentication/types"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (c *controller) Register(ctx context.Context, req *types.ReqRegister) (res *types.ResRegister, err error) {
	db := c.cfg.MysqlClient

	if strings.Contains(req.Name, " ") {
		return nil, fmt.Errorf("invalid request: format name not valid")
	}

	hash, err := c.utilities.HashString(req.Password)
	if err != nil {
		return nil, err
	}

	defaultRole := "borrower"

	if _, err := db.ExecContext(
		ctx,
		`
		insert into users (id, name, email, password_hash, role, created_at)
		values (?,?,?,?,?,?)
		`,
		uuid.NewString(), req.Name, req.Email, hash, defaultRole, time.Now(),
	); err != nil {
		return nil, err
	}

	return nil, nil
}