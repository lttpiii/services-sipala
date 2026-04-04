package controllers

import (
	"context"
	"fmt"
	"services-sipala/services/authentication/types"
	"time"
)

func (c *controller) RefreshToken(ctx context.Context, req *types.ReqRefreshToken) (res *types.ResRefreshToken, err error) {
	db := c.cfg.MysqlClient

	claim, err := c.utilities.ValidateJWT(req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	if claim.Type != "refresh_token" {
		return nil, fmt.Errorf("unauthorized: invalid token type %v", claim.Type)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var (
		userID, token string
	)
	if err := tx.QueryRowContext(
		ctx,
		`
		select
			user_id,
			token
		from refresh_tokens
		where id = ?
		and revoked_at is null
		and expires_at > ?
		`, claim.ID, time.Now(),
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

	accessTTL := 15 * time.Minute
	refreshTTL := 7 * 24 * time.Hour

	_, accessToken, err := c.utilities.GenerateJWT(userID, "access_token", claim.Role, accessTTL)
	if err != nil {
		return nil, err
	}

	refreshID, refreshToken, err := c.utilities.GenerateJWT(userID, "refresh_token", claim.Role, refreshTTL)
	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(
		ctx,
		`
		insert into refresh_tokens (id, user_id, token, expires_at, created_at)
		values (?,?,?,?,?)
		`,
		refreshID, userID, refreshToken, time.Now().Add(refreshTTL), time.Now(),
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &types.ResRefreshToken{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		TokenType: "bearer",
		ExpiresIn: int(accessTTL.Seconds()),
	}, nil
}