package controllers

import (
	"context"
	"fmt"
	"services-sipala/services/borrow/types"
	"time"
)

func (c *controller) GetBorrowByID(ctx context.Context, req *types.ReqGetBorrowByID) (res *types.ResGetBorrowByID, err error) {
	db := c.cfg.MysqlClient
	
	if req.AuthUserRole != "admin" && req.AuthUserRole != "staff" {
		var borrowerID string

		err := db.QueryRowContext(
			ctx,
			`
			SELECT borrower_id
			FROM borrow_transactions
			WHERE id = ?
			`,
			req.BorrowID,
		).Scan(&borrowerID)

		if err != nil {
			return nil, err
		}

		if borrowerID != req.AuthUserID {
			return nil, fmt.Errorf("forbidden: user must be owner borrow")
		}
	}

	var (
		result *types.ResGetBorrowByID
	)

	rows, err := db.QueryContext(
		ctx,
		`
		select
			bt.id,
			
			u.id,
			u.name,
			u.email,

			bt.status,
			bt.borrow_date,
			bt.due_date,

			bti.id,
			bti.tool_id,
			t.name,
			bti.quantity,

			bt.approved_by,
			bt.approved_at,
			bt.created_at

		from borrow_transactions bt
		join users u on u.id = bt.borrower_id
		left join borrow_transaction_items bti on bti.borrow_transaction_id = bt.id
		left join tools t on t.id = bti.tool_id
		where bt.id = ?
		`, req.BorrowID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			btID, btStatus string
			btBorrowDate, btDuedate, btCreatedAt time.Time
			btApprovedBy *string
			btApprovedAt *time.Time

			uID, uName, uEmail string

			itemID, toolID, toolName *string
			qty *int
		)

		if err := rows.Scan(
			&btID,

			&uID,
			&uName,
			&uEmail,

			&btStatus,
			&btBorrowDate,
			&btDuedate,

			&itemID,
			&toolID,
			&toolName,
			&qty,

			&btApprovedBy,
			&btApprovedAt,
			&btCreatedAt,
		); err != nil {
			return nil, err
		}

		if result == nil {
			result = &types.ResGetBorrowByID{
				ID: btID,
				Borrower: types.BorrowerType{
					ID: uID,
					Name: uName,
					Email: &uEmail,
				},
				Status: btStatus,
				BorrowDate: btBorrowDate,
				DueDate: btDuedate,
				Items: []types.ItemsType{},
				ApprovedBy: btApprovedBy,
				ApprovedAt: btApprovedAt,
				CreatedAt: btCreatedAt,
			}
		}

		if itemID != nil {
			item := types.ItemsType{
				ID: *itemID,
			}

			if toolID != nil {
				item.ToolID = *toolID
			}
			if toolName != nil {
				item.ToolName = *toolName
			}
			if qty != nil {
				item.Qty = *qty
			}

			result.Items = append(result.Items, item)
		}
	}

	if result == nil {
		return nil, fmt.Errorf("not found: borrow not found")
	}

	return result, nil
}