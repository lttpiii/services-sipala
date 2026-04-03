package controllers

import (
	"context"
	"services-sipala/services/borrow/types"
)

type (
	IBorrowController interface {
		CreateBorrow(ctx context.Context, req *types.ReqCreateBorrow) (res *types.ResCreateBorrow, err error)
		AddBorrowItem(ctx context.Context, req *types.ReqAddBorrowItem) (res *types.ResAddBorrowItem, err error)
		RemoveBorrowItem(ctx context.Context, req *types.ReqRemoveBorrowItem) (res *types.ResRemoveBorrowItem, err error)
		SubmitBorrow(ctx context.Context, req *types.ReqSubmitBorrow) (res *types.ResSubmitBorrow, err error)
		GetBorrowByID(ctx context.Context, req *types.ReqGetBorrowByID) (res *types.ResGetBorrowByID, err error)
		GetListBorrows(ctx context.Context, req *types.ReqGetListBorrows) (res *types.ResGetListBorrows, err error)
		GetMyBorrows(ctx context.Context, req *types.ReqGetMyBorrows) (res *types.ResGetMyBorrows, err error)
	}
)