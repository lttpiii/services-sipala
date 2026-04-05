package controllers

import (
	"context"
	"fmt"
	"math"
	"services-sipala/services/tools/types"

	"strings"
)

func (c *controller) GetListTools(ctx context.Context, req *types.ReqGetListTools) (res *types.ResGetListTools, err error) {
	db := c.cfg.MysqlClient

	var (
		tools []types.ToolType

		conds []string
		args []any
		order string
		totalRecords int
	)

	order = "ORDER BY t.created_at DESC"

	conds = append(conds, "t.deleted_at is null")

	if req.Search != "" {
		conds = append(conds, "t.name like ?")
		args = append(args, "%"+req.Search+"%")
	}

	if req.CategoryID != "" {
		conds = append(conds, "t.category_id = ?")
		args = append(args, req.CategoryID)
	}

	if req.AvailableOnly {
		conds = append(conds, "t.stock > 0")
	}

	whereClause := ""
	if len(conds) > 0 {
		whereClause = "WHERE " + strings.Join(conds, " AND ")
	}

	countArgs := append([]any{}, args...)

	if err := db.QueryRowContext(
		ctx,
		fmt.Sprintf(`
		select count(distinct t.id)
		from tools t
		join categories c on c.id = t.category_id
		left join borrow_transaction_items bti on bti.tool_id = t.id
		left join borrow_transactions bt 
		on bt.id = bti.borrow_transaction_id 
			and bt.status NOT IN ('returned', 'rejected')
		%s`, whereClause),
		countArgs...
	).Scan(&totalRecords); err != nil {
		return nil, err
	}

	offset := (req.Page - 1) * req.Limit
	dataArgs := append(args, req.Limit, offset)

	dataQuery := fmt.Sprintf(`
	select
		t.id,
		t.name,

		c.id,
		c.name,

		t.stock,
		t.stock - coalesce(SUM(bti.quantity), 0) as available_stock,
		t.description

	from tools t
	join categories c on c.id = t.category_id
	left join borrow_transaction_items bti on bti.tool_id = t.id
	left join borrow_transactions bt 
		on bt.id = bti.borrow_transaction_id 
		and bt.status NOT IN ('returned', 'rejected')
	%s
	group by
		t.id, t.name,
		c.id, c.name,
		t.stock,
		t.description,
		t.created_at,
		t.updated_at
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
		var (
			t types.ToolType
			ct types.CategoryType
		)

		if err := rows.Scan(
			&t.ID,
			&t.Name,
			&ct.ID,
			&ct.Name,
			&t.Stock,
			&t.AvailableStock,
			&t.Description,
		); err != nil {
			return nil, err
		}

		t.Category = ct

		tools = append(tools, t)
	}

	totalPages := int(math.Ceil(float64(totalRecords)/ float64(req.Limit)))

	return &types.ResGetListTools{
		Data: tools,
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