package controllers

import (
	"context"
	"fmt"
	"services-sipala/services/authentication/types"
	"time"
)

func (c *controller) Login(ctx context.Context, req *types.ReqLogin) (res *types.ResLogin, err error) {
	db := c.cfg.MysqlClient

	var (
		userID, username, email, role, passwordHash string
	)

	if err := db.QueryRowContext(
		ctx,
		`select
			id,
			name,
			email,
			role,
			password_hash
		from users
		where email = ?	
		`,
		req.Email,
	).Scan(
		&userID,
		&username,
		&email,
		&role,
		&passwordHash,
	); err != nil {
		return nil, err
	}

	if !c.utilities.CompareStringWithHash(req.Password, passwordHash) {
		return nil, fmt.Errorf("unauthorized: password not valid")
	}

	accessTTL := 15 * time.Minute
	refreshTTL := 7 * 24 * time.Hour

	_, accessToken, err := c.utilities.GenerateJWT(userID, "access_token", role, accessTTL)
	if err != nil {
		return nil, err
	}

	refreshID, refreshToken, err := c.utilities.GenerateJWT(userID, "refresh_token", role, refreshTTL)
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

	return &types.ResLogin{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		TokenType: "bearer",
		ExpiresIn: int(accessTTL.Seconds()),
		User: types.UserType{
			ID: userID,
			Name: username,
			Email: email,
			Role: role,
		},
	}, nil
}