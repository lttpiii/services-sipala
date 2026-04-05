package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"services-sipala/services/returns/types"
	"strings"
)

func (c *controller) GetListReturns(ctx context.Context, req *types.ReqGetListReturns) (res *types.ResGetListReturns, err error) {
	db := c.cfg.MysqlClient

	// =============================
	// 1. Default pagination
	// =============================
	page := req.Page
	if page <= 0 {
		page = 1
	}

	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit

	// =============================
	// 2. Dynamic WHERE builder
	// =============================
	var conditions []string
	var args []interface{}

	// filter borrower_id
	if req.BorrowerID != "" {
		conditions = append(conditions, "bt.borrower_id = ?")
		args = append(args, req.BorrowerID)
	}

	// filter date range
	if req.StartDate != nil {
		conditions = append(conditions, "rt.returned_at >= ?")
		args = append(args, *req.StartDate)
	}

	if req.EndDate != nil {
		conditions = append(conditions, "rt.returned_at <= ?")
		args = append(args, *req.EndDate)
	}

	// filter has_fine
	if req.HasFine {
		conditions = append(conditions, "rt.fine_amount > 0")
	}

	// search borrower name
	if req.Search != "" {
		conditions = append(conditions, "LOWER(u.name) LIKE ?")
		args = append(args, "%"+strings.ToLower(req.Search)+"%")
	}

	// build WHERE
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// =============================
	// 3. COUNT QUERY
	// =============================
	countQuery := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM return_transactions rt
		JOIN borrow_transactions bt ON bt.id = rt.borrow_transaction_id
		JOIN users u ON u.id = bt.borrower_id
		%s
	`, whereClause)

	var totalRecords int
	err = db.QueryRowContext(ctx, countQuery, args...).Scan(&totalRecords)
	if err != nil {
		return nil, err
	}

	// =============================
	// 4. DATA QUERY
	// =============================
	dataQuery := fmt.Sprintf(`
		SELECT
			rt.id,
			rt.borrow_transaction_id,
			u.name,
			rt.returned_at,
			rt.late_days,
			rt.fine_amount
		FROM return_transactions rt
		JOIN borrow_transactions bt ON bt.id = rt.borrow_transaction_id
		JOIN users u ON u.id = bt.borrower_id
		%s
		ORDER BY rt.returned_at DESC
		LIMIT ? OFFSET ?
	`, whereClause)

	dataArgs := append(args, limit, offset)

	rows, err := db.QueryContext(ctx, dataQuery, dataArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []types.Return

	for rows.Next() {
		var item types.Return

		var lateDays sql.NullInt64
		var fineAmount sql.NullFloat64
		var borrowerName string

		err := rows.Scan(
			&item.ID,
			&item.BorrowTransactionID,
			&borrowerName,
			&item.ReturnedAt,
			&lateDays,
			&fineAmount,
		)
		if err != nil {
			return nil, err
		}

		// map nullable
		if lateDays.Valid {
			val := int(lateDays.Int64)
			item.LateDays = &val
		}
		if fineAmount.Valid {
			val := float64(fineAmount.Float64)
			item.FineAmount = &val
		}

		// inject borrower name ke BorrowDetails (biar konsisten struct kamu)
		item.BorrowDetails.BorrowerName = borrowerName

		data = append(data, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// =============================
	// 5. Metadata
	// =============================
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	res = &types.ResGetListReturns{
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