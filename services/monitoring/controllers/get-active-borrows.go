package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"strings"
	"time"

	"services-sipala/services/monitoring/types"
)

func (c *controller) GetActiveBorrows(ctx context.Context, req *types.ReqGetActiveBorrows) (res *types.ResGetActiveBorrows, err error) {
	db := c.cfg.MysqlClient

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
	search := strings.TrimSpace(req.Search)
	searchLike := "%" + search + "%"


	countQuery := `
		SELECT COUNT(*)
		FROM borrow_transactions bt
		INNER JOIN users u ON u.id = bt.borrower_id
		LEFT JOIN return_transactions rt ON rt.borrow_transaction_id = bt.id
		WHERE bt.status = ?
		  AND rt.id IS NULL
		  AND u.deleted_at IS NULL
		  AND (? = '' OR u.name LIKE ?)
	`

	var totalRecords int
	err = db.QueryRowContext(
		ctx,
		countQuery,
		"approved",
		search,
		searchLike,
	).Scan(&totalRecords)
	if err != nil {
		return nil, fmt.Errorf("failed count active borrows: %w", err)
	}

	mainQuery := `
		SELECT 
			bt.id,
			bt.borrower_id,
			u.name,
			u.email,
			bt.borrow_date,
			bt.due_date,
			COALESCE(SUM(bti.quantity), 0) AS total_items
		FROM borrow_transactions bt
		INNER JOIN users u ON u.id = bt.borrower_id
		LEFT JOIN return_transactions rt ON rt.borrow_transaction_id = bt.id
		LEFT JOIN borrow_transaction_items bti ON bti.borrow_transaction_id = bt.id
		WHERE bt.status = ?
		  AND rt.id IS NULL
		  AND u.deleted_at IS NULL
		  AND (? = '' OR u.name LIKE ?)
		GROUP BY 
			bt.id,
			bt.borrower_id,
			u.name,
			u.email,
			bt.borrow_date,
			bt.due_date
		ORDER BY bt.due_date ASC
		LIMIT ? OFFSET ?
	`

	rows, err := db.QueryContext(
		ctx,
		mainQuery,
		"approved",
		search,
		searchLike,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("failed query active borrows: %w", err)
	}
	defer rows.Close()

	data := make([]types.BorrowType, 0)
borrowIDs := make([]string, 0)
	now := time.Now()

	for rows.Next() {
		var item types.BorrowType
		var borrowerID string
		var borrowerName string
		var borrowerEmail sql.NullString

		err = rows.Scan(
			&item.ID,
			&borrowerID,
			&borrowerName,
			&borrowerEmail,
			&item.BorrowDate,
			&item.DueDate,
			&item.TotalItems,
		)
		if err != nil {
			return nil, fmt.Errorf("failed scan active borrows: %w", err)
		}

		var emailPtr *string
		if borrowerEmail.Valid {
			email := borrowerEmail.String
			emailPtr = &email
		}

		item.Borrower = &types.BorrowerType{
			ID:    borrowerID,
			Name:  borrowerName,
			Email: emailPtr,
		}


		daysRemaining := int(math.Ceil(item.DueDate.Sub(now).Hours() / 24))
		isOverdue := now.After(item.DueDate)

		item.DaysRemaining = daysRemaining
		item.IsOverdue = isOverdue
		item.Items = []types.ItemsType{}

		data = append(data, item)
		borrowIDs = append(borrowIDs, item.ID)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed iterate active borrows: %w", err)
	}


	if len(borrowIDs) > 0 {
		placeholders := make([]string, len(borrowIDs))
		args := make([]interface{}, len(borrowIDs))

		for i, id := range borrowIDs {
			placeholders[i] = "?"
			args[i] = id
		}

		itemsQuery := fmt.Sprintf(`
			SELECT 
				bti.borrow_transaction_id,
				t.name,
				bti.quantity
			FROM borrow_transaction_items bti
			INNER JOIN tools t ON t.id = bti.tool_id
			WHERE bti.borrow_transaction_id IN (%s)
			  AND t.deleted_at IS NULL
			ORDER BY t.name ASC
		`, strings.Join(placeholders, ","))

		itemRows, err := db.QueryContext(ctx, itemsQuery, args...)
		if err != nil {
			return nil, fmt.Errorf("failed query borrow items: %w", err)
		}
		defer itemRows.Close()

		itemsMap := make(map[string][]types.ItemsType)

		for itemRows.Next() {
			var borrowID string
			var toolName string
			var qty int

			err = itemRows.Scan(&borrowID, &toolName, &qty)
			if err != nil {
				return nil, fmt.Errorf("failed scan borrow items: %w", err)
			}

			itemsMap[borrowID] = append(itemsMap[borrowID], types.ItemsType{
				ToolName: toolName,
				Qty:      qty,
			})
		}

		if err = itemRows.Err(); err != nil {
			return nil, fmt.Errorf("failed iterate borrow items: %w", err)
		}

		for i := range data {
			if items, ok := itemsMap[data[i].ID]; ok {
				data[i].Items = items
			}
		}
	}


	totalPages := 0
	if totalRecords > 0 {
		totalPages = int(math.Ceil(float64(totalRecords) / float64(limit)))
	}

	res = &types.ResGetActiveBorrows{
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