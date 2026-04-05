package controllers

import (
	"context"
	"services-sipala/services/logs/types"
)
type (
	ILogsController interface {
		GetLogs(ctx context.Context, req *types.ReqGetLogs) (res *types.ResGetLogs, err error)
	}
)