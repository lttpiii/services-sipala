package controllers

import (
	"context"
	"services-sipala/services/categories/types"
)

func (c *controller) GetCategoryByID(ctx context.Context, req *types.ReqGetCategoryByID) (res *types.ResGetCategoryByID, err error) {
	db := c.cfg.MysqlClient

	var category types.CategoryType

	if err := db.QueryRowContext(
		ctx,
		`
		select
			c.id,
			c.name,
			c.created_at,
			c.updated_at,

			coalesce(count(t.id)) as total_tools
		from categories c
		left join tools t on t.category_id = c.id
		where c.id = ?
		group by c.id, c.name
		`, req.CategoryID,
	).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.ToolsCount,
	); err != nil {
		return nil, err
	}

	return &types.ResGetCategoryByID{
		ID: category.ID,
		Name: category.Name,
		ToolsCount: category.ToolsCount,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}