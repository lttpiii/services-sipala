package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"services-sipala/services/reporting/types"
)

func (c *controller) GetBorrowReport(ctx context.Context, req *types.ReqGetBorrowReport) (res *types.ResGetBorrowReport, err error) {
	db := c.cfg.MysqlClient

	if req.GroupBy == "" {
		req.GroupBy = "day"
	}

	req.GroupBy = strings.ToLower(req.GroupBy)

	if req.GroupBy != "day" && req.GroupBy != "week" && req.GroupBy != "month" {
		return nil, fmt.Errorf("invalid group_by, allowed values: day, week, month")
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start_date format, expected YYYY-MM-DD")
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end_date format, expected YYYY-MM-DD")
	}

	if endDate.Before(startDate) {
		return nil, fmt.Errorf("end_date must be greater than or equal to start_date")
	}


	endDateInclusive := endDate.Add(24 * time.Hour)

	var summary types.BorrowReportSummaryType


	summaryQuery := `
		SELECT
			COUNT(DISTINCT bt.id) AS total_transactions,
			COALESCE(SUM(bti.quantity), 0) AS total_items_borrowed,
			COALESCE(COUNT(DISTINCT CASE WHEN bt.status = 'approved' THEN bt.id END), 0) AS approved,
			COALESCE(COUNT(DISTINCT CASE WHEN bt.status = 'rejected' THEN bt.id END), 0) AS rejected,
			COALESCE(COUNT(DISTINCT CASE WHEN bt.status = 'pending' THEN bt.id END), 0) AS pending
		FROM borrow_transactions bt
		LEFT JOIN borrow_transaction_items bti ON bti.borrow_transaction_id = bt.id
		WHERE bt.borrow_date >= ? AND bt.borrow_date < ?
	`

	if err := db.QueryRowContext(
		ctx,
		summaryQuery,
		startDate,
		endDateInclusive,
	).Scan(
		&summary.TotalTransactions,
		&summary.TotalItemsBorrowed,
		&summary.Approved,
		&summary.Rejected,
		&summary.Pending,
	); err != nil {
		return nil, err
	}

	topTools := []types.BorrowReportTopToolType{}

	topToolsQuery := `
		SELECT
			t.id AS tool_id,
			t.name AS tool_name,
			COUNT(DISTINCT bt.id) AS borrow_count,
			COALESCE(SUM(bti.quantity), 0) AS total_quantity
		FROM borrow_transactions bt
		JOIN borrow_transaction_items bti ON bti.borrow_transaction_id = bt.id
		JOIN tools t ON t.id = bti.tool_id
		WHERE bt.borrow_date >= ? AND bt.borrow_date < ?
		GROUP BY t.id, t.name
		ORDER BY borrow_count DESC, total_quantity DESC, t.name ASC
		LIMIT 10
	`

	topRows, err := db.QueryContext(
		ctx,
		topToolsQuery,
		startDate,
		endDateInclusive,
	)
	if err != nil {
		return nil, err
	}
	defer topRows.Close()

	for topRows.Next() {
		var item types.BorrowReportTopToolType

		if err := topRows.Scan(
			&item.ToolID,
			&item.ToolName,
			&item.BorrowCount,
			&item.TotalQuantity,
		); err != nil {
			return nil, err
		}

		topTools = append(topTools, item)
	}


	byDate := []types.BorrowReportByDateType{}

	var groupExpr string
	switch req.GroupBy {
	case "day":
		groupExpr = "DATE(bt.borrow_date)"
	case "week":

		groupExpr = "DATE_FORMAT(bt.borrow_date, '%x-W%v')"
	case "month":
		groupExpr = "DATE_FORMAT(bt.borrow_date, '%Y-%m')"
	}

	byDateQuery := fmt.Sprintf(`
		SELECT
			%s AS grouped_date,
			COUNT(DISTINCT bt.id) AS transaction_count,
			COALESCE(SUM(bti.quantity), 0) AS item_count
		FROM borrow_transactions bt
		LEFT JOIN borrow_transaction_items bti ON bti.borrow_transaction_id = bt.id
		WHERE bt.borrow_date >= ? AND bt.borrow_date < ?
		GROUP BY grouped_date
		ORDER BY grouped_date ASC
	`, groupExpr)

	byDateRows, err := db.QueryContext(
		ctx,
		byDateQuery,
		startDate,
		endDateInclusive,
	)
	if err != nil {
		return nil, err
	}
	defer byDateRows.Close()

	for byDateRows.Next() {
		var (
			item      types.BorrowReportByDateType
			groupedAt sql.NullString
		)

		if err := byDateRows.Scan(
			&groupedAt,
			&item.TransactionCount,
			&item.ItemCount,
		); err != nil {
			return nil, err
		}

		if groupedAt.Valid {
			item.Date = groupedAt.String
		}

		byDate = append(byDate, item)
	}

	return &types.ResGetBorrowReport{
		Period: types.BorrowReportPeriodType{
			StartDate: req.StartDate,
			EndDate:   req.EndDate,
		},
		Summary:  summary,
		TopTools: topTools,
		ByDate:   byDate,
	}, nil
}