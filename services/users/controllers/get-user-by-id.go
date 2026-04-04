package controllers

import (
	"context"
	"fmt"
	"services-sipala/services/users/types"
	"strings"
)

func (c *controller) GetUserByID(ctx context.Context, req *types.ReqGetUserByID) (res *types.ResGetUserByID, err error) {
	db := c.cfg.MysqlClient

	var (
		conds []string
		args []any
	)

	if req.AuthUserRole == "admin" {
		if !req.IncludeDeleted {
			conds = append(conds, "deleted_at is null")
		}
	} else {
		conds = append(conds, "deleted_at is null")
	}

	conds = append(conds, "id = ?")
	args = append(args, req.UserID)

	whereClause := ""
	if len(conds) > 0 {
		whereClause = "WHERE " + strings.Join(conds, " AND ")
	}

	query := fmt.Sprintf(`
		select
			id,
			name,
			email,
			role,
			created_at,
			updated_at,
			deleted_at
		from users
		%s
	`, whereClause)

	var u types.UserType

	if err := db.QueryRowContext(
		ctx,
		query,
		args...
	).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	); err != nil {
		return nil, err
	}

	return &types.ResGetUserByID{
		ID: u.ID,
		Name: u.Name,
		Email: u.Email,
		Role: u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}, nil
}