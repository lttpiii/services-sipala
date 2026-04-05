package controllers

import (
	"context"
	"fmt"
	"math"
	"services-sipala/services/borrow/types"
	"strings"
)

func (c *controller) GetMyBorrows(ctx context.Context, req *types.ReqGetMyBorrows) (res *types.ResGetMyBorrows, err error) {
	db := c.cfg.MysqlClient

	var (
		results      []types.BorrowType
		conds        []string
		args         []any
		totalRecords int
	)

	// ✅ WAJIB: selalu masukkan borrower_id ke conds
	conds = append(conds, "bt.borrower_id = ?")
	args = append(args, req.UserID)

	if req.Status != "" {
		conds = append(conds, "bt.status = ?")
		args = append(args, req.Status)
	}

	// build WHERE sekali saja
	whereClause := ""
	if len(conds) > 0 {
		whereClause = "WHERE " + strings.Join(conds, " AND ")
	}

	// ======================
	// 🔢 COUNT QUERY
	// ======================
	countQuery := fmt.Sprintf(`
		SELECT COUNT(DISTINCT bt.id)
		FROM borrow_transactions bt
		LEFT JOIN borrow_transaction_items bti 
			ON bti.borrow_transaction_id = bt.id
		%s
	`, whereClause)

	if err := db.QueryRowContext(
		ctx,
		countQuery,
		args...,
	).Scan(&totalRecords); err != nil {
		return nil, err
	}

	// ======================
	// 📄 DATA QUERY
	// ======================
	offset := (req.Page - 1) * req.Limit

	dataQuery := fmt.Sprintf(`
		SELECT
			bt.id,
			bt.status,
			bt.borrow_date,
			bt.due_date,

			COUNT(DISTINCT bti.id) as total_items,
			COALESCE(SUM(bti.quantity), 0) as total_quantity,

			CASE 
				WHEN bt.due_date < NOW() AND bt.status NOT IN ('returned', 'rejected')
				THEN TRUE
				ELSE FALSE
			END as is_overdue

		FROM borrow_transactions bt
		LEFT JOIN borrow_transaction_items bti 
			ON bti.borrow_transaction_id = bt.id

		%s

		GROUP BY 
			bt.id, bt.status, bt.borrow_date, bt.due_date

		ORDER BY bt.created_at DESC
		LIMIT ? OFFSET ?
	`, whereClause)

	dataArgs := append(args, req.Limit, offset)

	rows, err := db.QueryContext(
		ctx,
		dataQuery,
		dataArgs...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {	
		var item types.BorrowType

		if err := rows.Scan(
			&item.ID,
			&item.Status,
			&item.BorrowDate,
			&item.DueDate,
			&item.TotalItems,
			&item.TotalQuantity,
			&item.IsOverdue,
		); err != nil {
			return nil, err
		}

		results = append(results, item)
	}

	// optional tapi best practice
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// ======================
	// 📊 PAGINATION
	// ======================
	totalPages := 0
	if req.Limit > 0 {
		totalPages = int(math.Ceil(float64(totalRecords) / float64(req.Limit)))
	}

	return &types.ResGetMyBorrows{
		Data: results,
		Metadata: types.MetadataType{
			CurrentPage:  req.Page,
			PageSize:     req.Limit,
			TotalPages:   totalPages,
			TotalRecords: totalRecords,
			HasNext:      req.Page < totalPages,
			HasPrev:      req.Page > 1,
		},
	}, nil
}