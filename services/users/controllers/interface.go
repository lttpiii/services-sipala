package controllers

import (
	"context"
	"services-sipala/services/users/types"
)

type (
	IUsersController interface {
		CreateUser(ctx context.Context, req *types.ReqCreateUser) (res *types.ResCreateUser, err error)
		UpdateUser(ctx context.Context, req *types.ReqUpdateUser) (res *types.ResUpdateUser, err error)
		DeleteUser(ctx context.Context, req *types.ReqDeleteUser) (res *types.ResDeleteUser, err error)
		GetUserByID(ctx context.Context, req *types.ReqGetUserByID) (res *types.ResGetUserByID, err error)
		GetListUsers(ctx context.Context, req *types.ReqGetListUsers) (res *types.ResGetListUsers, err error)
	}
)