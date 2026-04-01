package controllers

import (
	"context"
	"services-sipala/services/authentication/types"
)

type (
	IAuthenticationController interface {
		Login(ctx context.Context, req *types.ReqLogin) (res *types.ResLogin, err error)
		Logout(ctx context.Context, req *types.ReqLogout) (res *types.ResLogout, err error)
		Register(ctx context.Context, req *types.ReqRegister) (res *types.ResRegister, err error)
		RefreshToken(ctx context.Context, req *types.ReqRefreshToken) (res *types.ResRefreshToken, err error)
		GetProfile(ctx context.Context, req *types.ReqGetProfile) (res *types.ResGetProfile, err error)
		ChangePassword(ctx context.Context, req *types.ReqChangePassword) (res *types.ResChangePassword, err error)
	}
)