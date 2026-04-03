package controllers

import (
	"context"
	"services-sipala/services/tools/types"
)

type (
	IToolsController interface {
		CreateTools(ctx context.Context, req *types.ReqCreateTool) (res *types.ReqCreateTool, err error)
		UpdateTools(ctx context.Context, req *types.ReqUpdateTool) (res *types.ReqUpdateTool, err error)
		DeleteTools(ctx context.Context, req *types.ReqDeleteTool) (res *types.ReqDeleteTool, err error)
		GetToolsByID(ctx context.Context, req *types.ReqGetToolByID) (res *types.ReqGetToolByID, err error)
		ListTools(ctx context.Context, req *types.ReqListTools) (res *types.ReqListTools, err error)
	}
)
