package controllers

import (
	"context"
	"services-sipala/services/authentication/types"
)

func (c *controller) GetProfile(ctx context.Context, req *types.ReqGetProfile) (res *types.ResGetProfile, err error) {
	db := c.cfg.MysqlClient

	var user types.ResGetProfile

	if err := db.QueryRowContext(
		ctx,
		`
		select
			id,
			name,
			email,
			role,
			created_at,
			updated_at
		from users
		where id = ?
		and deleted_at is null
		`, req.UserID,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}