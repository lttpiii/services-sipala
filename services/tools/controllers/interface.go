package controllers

import (
	"context"
	"services-sipala/services/tools/types"
)

type (
	IToolsController interface {
		CreateTool(ctx context.Context, req *types.ReqCreateTool) (res *types.ResCreateTool, err error)
		UpdateTool(ctx context.Context, req *types.ReqUpdateTool) (res *types.ResUpdateTool, err error)
		DeleteTool(ctx context.Context, req *types.ReqDeleteTool) (res *types.ResDeleteTool, err error)
		GetToolByID(ctx context.Context, req *types.ReqGetToolByID) (res *types.ResGetToolByID, err error)
		GetListTools(ctx context.Context, req *types.ReqGetListTools) (res *types.ResGetListTools, err error)
	}
)