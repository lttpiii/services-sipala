package controllers

import (
	"context"
	"fmt"
	"math"
	"services-sipala/services/borrow/types"
	"strings"
	"time"
)

func (c *controller) GetListBorrows(ctx context.Context, req *types.ReqGetListBorrows) (res *types.ResGetListBorrows, err error) {
	db := c.cfg.MysqlClient

	var (
		results []types.BorrowType

		conds []string
		args  []any

		totalRecords int
		order        = "ORDER BY bt.created_at DESC"
	)

	if req.Search != "" {
		conds = append(conds, "(u.name LIKE ? OR u.email LIKE ?)")
		args = append(args, "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if req.Status != "" {
		conds = append(conds, "bt.status = ?")
		args = append(args, req.Status)
	}

	if req.BorrowerID != "" {
		conds = append(conds, "u.id = ?")
		args = append(args, req.BorrowerID)
	}

	if req.StartDate != nil {
		conds = append(conds, "bt.borrow_date > ?")
		args = append(args, req.StartDate)	
	}

	if req.EndDate != nil {
		conds = append(conds, "bt.due_date < ?")
		args = append(args, req.EndDate)
	}

	whereClause := ""
	if len(conds) > 0 {
		whereClause = "WHERE " + strings.Join(conds, " AND ")
	}

	countArgs := append([]any{}, args...)
	if err := db.QueryRowContext(
		ctx,
		fmt.Sprintf(`
			SELECT COUNT(DISTINCT bt.id)
			FROM borrow_transactions bt
			JOIN users u ON u.id = bt.borrower_id
			LEFT JOIN borrow_transaction_items bti 
				ON bti.borrow_transaction_id = bt.id
			%s
		`, whereClause),
		countArgs...,
	).Scan(&totalRecords); err != nil {
		return nil, err
	}

	offset := (req.Page - 1) * req.Limit
	dataArgs := append(args, req.Limit, offset)

	rows, err := db.QueryContext(
		ctx,
		fmt.Sprintf(`
		SELECT
			bt.id,

			u.id,
			u.name,

			bt.status,
			bt.borrow_date,
			bt.due_date,

			COUNT(DISTINCT bti.id) as total_items,
			COALESCE(SUM(bti.quantity), 0) as total_quantity,

			bt.created_at
		FROM borrow_transactions bt
		JOIN users u ON u.id = bt.borrower_id
		LEFT JOIN borrow_transaction_items bti 
			ON bti.borrow_transaction_id = bt.id
		%s
		GROUP BY 
			bt.id, u.id, u.name,
			bt.status, bt.borrow_date, bt.due_date, bt.created_at
		%s
		LIMIT ? OFFSET ?
		`, whereClause, order),
		dataArgs...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item types.BorrowType
		item.Borrower = &types.BorrowerType{} // ✅ FIX


		if err := rows.Scan(
			&item.ID,

			&item.Borrower.ID,
			&item.Borrower.Name,

			&item.Status,
			&item.BorrowDate,
			&item.DueDate,

			&item.TotalItems,
			&item.TotalQuantity,

			new(time.Time), // skip created_at kalau gak dipakai
		); err != nil {
			return nil, err
		}

		results = append(results, item)
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(req.Limit)))

	return &types.ResGetListBorrows{
		Data: results,
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