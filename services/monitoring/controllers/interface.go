package controllers

import (
	"context"
	"services-sipala/services/monitoring/types"
)

type (
	IMonitoringController interface {
		GetActiveBorrows(ctx context.Context, req *types.ReqGetActiveBorrows) (res *types.ResGetActiveBorrows, err error)
		GetOverdueBorrows(ctx context.Context, req *types.ReqGetOverdueBorrows) (res *types.ResGetOverdueBorrows, err error)
	}
)