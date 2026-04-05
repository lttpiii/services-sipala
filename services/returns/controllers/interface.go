package controllers

import (
	"context"
	"services-sipala/services/returns/types"
)

type (
	IReturnController interface {
		CreateReturns(ctx context.Context, req *types.ReqCreateReturns) (res *types.ResCreateReturns, err error)
		CalculateFine(ctx context.Context, req *types.ReqCalculateFineReturns) (res *types.ResCalculateFineReturns, err error)
		GetReturnByID(ctx context.Context, req *types.ReqGetReturnByID) (res *types.ResGetReturnByID, err error)
		GetListReturns(ctx context.Context, req *types.ReqGetListReturns) (res *types.ResGetListReturns, err error)
	}
)