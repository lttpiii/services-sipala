package controllers

import (
	"context"
	"services-sipala/services/tools/types"
)

func (c *controller) GetToolByID(ctx context.Context, req *types.ReqGetToolByID) (res *types.ResGetToolByID, err error) {
	db := c.cfg.MysqlClient

	var ( 
		t types.ResGetToolByID
		ct types.CategoryType
	)

	if err := db.QueryRowContext(
		ctx,
		`
		select
			t.id,
			t.name,

			c.id,
			c.name,

			t.stock,
			t.stock - coalesce(SUM(bti.quantity), 0) as available_stock,
			count(bti.id),
			t.description,
			t.created_at,
			t.updated_at

		from tools t
		join categories c on c.id = t.category_id
		left join borrow_transaction_items bti on bti.tool_id = t.id
		left join borrow_transactions bt on bt.id = bti.borrow_transaction_id where ((bt.status <> 'returned' AND bt.status <> 'rejected')OR bt.id IS NULL)
		where t.id = ? and t.deleted_at is null
		group by
			t.id, t.name,
			c.id, c.name,
			t.stock,
			t.description,
			t.created_at,
			t.updated_at
		`, req.ToolID,
	).Scan(
		&t.ID,
		&t.Name,
		&ct.ID,
		&ct.Name,
		&t.Stock,
		&t.AvailableStock,
		&t.BorrowedCount,
		&t.Description,
		&t.CreatedAt,
		&t.UpdatedAt,
	); err != nil {
		return nil, err
	}

	t.Category = ct

	return &t, nil
}