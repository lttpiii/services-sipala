package controllers

import (
	"context"
	"services-sipala/services/categories/types"
)

type (
	ICategoriesController interface {
		CreateCategory(ctx context.Context, req *types.ReqCreateCategory) (res *types.ResCreateCategory, err error)
		UpdateCategory(ctx context.Context, req *types.ReqUpdateCategory) (res *types.ResUpdateCategory, err error)
		DeleteCategory(ctx context.Context, req *types.ReqDeleteCategory) (res *types.ResDeleteCategory, err error)
		GetCategoryByID(ctx context.Context, req *types.ReqGetCategoryByID) (res *types.ResGetCategoryByID, err error)
		GetListCategories(ctx context.Context, req *types.ReqGetListCategories) (res *types.ResGetListCategories, err error)
	}
)