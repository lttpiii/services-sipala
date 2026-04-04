package controllers

import (
	"context"
	"fmt"
	"math"
	"services-sipala/services/categories/types"
	"strings"
)

func (c *controller) GetListCategories(ctx context.Context, req *types.ReqGetListCategories) (res *types.ResGetListCategories, err error) {
	db := c.cfg.MysqlClient

	var (
		categories []types.CategoryType

		conds []string
		args []any
		order string
		totalRecords int
	)

	order = "ORDER BY c.created_at DESC"

	if req.Search != "" {
		conds = append(conds, "c.name like ?")
		args = append(args, "%"+req.Search+"%")
	}

	whereClause := ""
	if len(conds) > 0 {
		whereClause = "WHERE " + strings.Join(conds, " AND ")
	}

	countArgs := append([]any{}, args...)

	if err := db.QueryRowContext(
		ctx,
		fmt.Sprintf("select count(*) from categories c %s", whereClause),
		countArgs...
	).Scan(&totalRecords); err != nil {
		return nil, err
	}

	offset := (req.Page - 1) * req.Limit
	dataArgs := append(args, req.Limit, offset)

	dataQuery := fmt.Sprintf(`
	select
		c.id,
		c.name,

		count(t.id) as total_tools
	from categories c
	left join tools t on t.category_id = c.id
	%s
	group by c.id, c.name
	%s
	limit ? offset ?
	`, whereClause, order)

	rows, err := db.QueryContext(
		ctx, dataQuery, dataArgs...
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category types.CategoryType

		if err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.ToolsCount,
		); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	totalPages := int(math.Ceil(float64(totalRecords)/ float64(req.Limit)))

	return &types.ResGetListCategories{
		Data: categories,
		Metadata: types.MetadataType{
			CurrentPage: req.Page,
			PageSize: req.Limit,
			TotalPages: totalPages,
			TotalRecords: totalRecords,
			HasNext: req.Page < totalPages,
			HasPrev: req.Page > 1,
		},
	}, nil
}