package controllers

import (
	"context"
	"fmt"
	"services-sipala/services/authentication/types"
	"time"
)

func (c *controller) Logout(ctx context.Context, req *types.ReqLogout) (res *types.ResLogout, err error) {
	db := c.cfg.MysqlClient

	claim, err := c.utilities.ValidateJWT(req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	if claim.Type != "refresh_token" {
		return nil, fmt.Errorf("unauthorized: failed token type '%v'", claim.Type)
	}

	var (
		token, userID string
	)

	if err := db.QueryRowContext(
		ctx,
		`
		select
			user_id,
			token
		from refresh_tokens
		where id = ?
		and revoked_at is null
		and expires_at > ?
		`,
		claim.ID, time.Now(),
	).Scan(&userID, &token); err != nil {
		return nil, err
	}

	if userID != claim.Subject {
		return nil, fmt.Errorf("unauthorized: failed user id")
	}

	if token != req.RefreshToken {
		return nil, fmt.Errorf("unauthorized: invalid token")
	}

	result, err := db.ExecContext(
		ctx,
		`UPDATE refresh_tokens 
		 SET revoked_at = ? 
		 WHERE id = ? AND revoked_at IS NULL`,
		time.Now(), claim.ID,
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

	return nil, nil
}