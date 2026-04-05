package controllers

import (
	"context"
	"services-sipala/services/reporting/types"
)

type (
	IReportingController interface {
		GetBorrowReport(ctx context.Context, req *types.ReqGetBorrowReport) (res *types.ResGetBorrowReport, err error)
		GetReturnReport(ctx context.Context, req *types.ReqGetReturnReport) (res *types.ReqGetReturnReport, err error)
		GetFineReport(ctx context.Context, req *types.ReqGetFineReport) (res *types.ReqGetFineReport, err error)
	}
)
