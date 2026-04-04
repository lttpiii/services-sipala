package controllers

import (
	"context"
	"fmt"
	"services-sipala/services/authentication/types"
	"time"
)

func (c *controller) ChangePassword(ctx context.Context, req *types.ReqChangePassword) (res *types.ResChangePassword, err error) {
	db := c.cfg.MysqlClient

	var oldHashPassword string

	if err := db.QueryRowContext(
		ctx,
		`
		select 
			password_hash
		from users
		where id = ?
		and deleted_at is null	
		`, req.UserID,
	).Scan(&oldHashPassword); err != nil {
		return nil, err
	}

	if !c.utilities.CompareStringWithHash(req.OldPassword, oldHashPassword) {
		return nil, fmt.Errorf("invalid request: mismatch old password")
	}

	if req.NewPassword != req.ConfirmPassword {
		return nil, fmt.Errorf("invalid request: mismatch new password and confirm password")
	}

	hash, err := c.utilities.HashString(req.NewPassword)
	if err != nil {
		return nil, err
	}

	result, err := db.ExecContext(
		ctx,
		`
		update users
		set password_hash = ?, updated_at = ?
		where id = ? and deleted_at is null
		`,
		hash, time.Now(), req.UserID,
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

	return &types.ResChangePassword{}, nil
}