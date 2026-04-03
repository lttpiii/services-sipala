package controllers

import (
	"context"
	"services-sipala/services/borrow-approval/types"
)

type (
	IBorrowApprovalController interface {
		ApproveBorrow(ctx context.Context, req *types.ReqApproveBorrow) (res *types.ResApproveBorrow, err error)
		RejectBorrow(ctx context.Context, req *types.ReqRejectBorrow) (res *types.ResRejectBorrow, err error)
	}
)