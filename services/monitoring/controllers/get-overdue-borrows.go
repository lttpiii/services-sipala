package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"services-sipala/services/monitoring/types"
)

func (c *controller) GetOverdueBorrows(ctx context.Context, req *types.ReqGetOverdueBorrows) (res *types.ResGetOverdueBorrows, err error) {
	db := c.cfg.MysqlClient

	page := req.Page
	limit := req.Limit

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit

	const finePerDay = 10000

	countQuery := `
		SELECT COUNT(DISTINCT bt.id)
		FROM borrow_transactions bt
		JOIN users u ON u.id = bt.borrower_id
		LEFT JOIN return_transactions rt ON rt.borrow_transaction_id = bt.id
		WHERE bt.status = 'approved'
		  AND bt.due_date < NOW()
		  AND rt.id IS NULL
		  AND u.deleted_at IS NULL
	`

	var countArgs []interface{}

	if strings.TrimSpace(req.Search) != "" {
		countQuery += " AND LOWER(u.name) LIKE ?"
		countArgs = append(countArgs, "%"+strings.ToLower(strings.TrimSpace(req.Search))+"%")
	}

	var totalRecords int
	err = db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&totalRecords)
	if err != nil {
		return nil, fmt.Errorf("failed count overdue borrows: %w", err)
	}

	mainQuery := `
		SELECT 
			bt.id,
			u.id,
			u.name,
			u.email,
			bt.borrow_date,
			bt.due_date,
			DATEDIFF(NOW(), bt.due_date) AS late_days
		FROM borrow_transactions bt
		JOIN users u ON u.id = bt.borrower_id
		LEFT JOIN return_transactions rt ON rt.borrow_transaction_id = bt.id
		WHERE bt.status = 'approved'
		  AND bt.due_date < NOW()
		  AND rt.id IS NULL
		  AND u.deleted_at IS NULL
	`

	var args []interface{}

	if strings.TrimSpace(req.Search) != "" {
		mainQuery += " AND LOWER(u.name) LIKE ?"
		args = append(args, "%"+strings.ToLower(strings.TrimSpace(req.Search))+"%")
	}

	mainQuery += `
		ORDER BY bt.due_date ASC
		LIMIT ? OFFSET ?
	`
	args = append(args, limit, offset)

	rows, err := db.QueryContext(ctx, mainQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed query overdue borrows: %w", err)
	}
	defer rows.Close()

	var data []types.OverdueBorrowType

	for rows.Next() {
		var item types.OverdueBorrowType
		var borrower types.BorrowerType
		var email sql.NullString

		err = rows.Scan(
			&item.ID,
			&borrower.ID,
			&borrower.Name,
			&email,
			&item.BorrowDate,
			&item.DueDate,
			&item.LateDays,
		)
		if err != nil {
			return nil, fmt.Errorf("failed scan overdue borrow: %w", err)
		}

		if email.Valid {
			borrower.Email = &email.String
		}

		item.Borrower = &borrower
		item.EstimatedFine = item.LateDays * finePerDay

		itemQuery := `
			SELECT 
				t.name,
				bti.quantity
			FROM borrow_transaction_items bti
			JOIN tools t ON t.id = bti.tool_id
			WHERE bti.borrow_transaction_id = ?
			  AND t.deleted_at IS NULL
		`

		itemRows, err := db.QueryContext(ctx, itemQuery, item.ID)
		if err != nil {
			return nil, fmt.Errorf("failed query overdue borrow items: %w", err)
		}

		var items []types.ItemsType
		totalItems := 0

		for itemRows.Next() {
			var detail types.ItemsType

			err = itemRows.Scan(
				&detail.ToolName,
				&detail.Qty,
			)
			if err != nil {
				itemRows.Close()
				return nil, fmt.Errorf("failed scan overdue borrow item: %w", err)
			}

			totalItems += detail.Qty
			items = append(items, detail)
		}

		if err = itemRows.Err(); err != nil {
			itemRows.Close()
			return nil, fmt.Errorf("error iterate overdue borrow items: %w", err)
		}

		itemRows.Close()

		item.Items = items
		item.TotalItems = totalItems

		data = append(data, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterate overdue borrows: %w", err)
	}

	totalPages := 0
	if totalRecords > 0 {
		totalPages = (totalRecords + limit - 1) / limit
	}

	res = &types.ResGetOverdueBorrows{
		Data: data,
		Metadata: types.MetadataType{
			CurrentPage:  page,
			PageSize:     limit,
			TotalPages:   totalPages,
			TotalRecords: totalRecords,
			HasNext:      page < totalPages,
			HasPrev:      page > 1,
		},
	}

	return res, nil
}