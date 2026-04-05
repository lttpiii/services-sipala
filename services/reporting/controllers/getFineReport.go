package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"services-sipala/services/reporting/types"
)

func (c *controller) GetFineReport(ctx context.Context, req *types.ReqGetFineReport) (res *types.ResGetFineReport, err error) {
	db := c.cfg.MysqlClient

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

	
	var (
		summary     types.FineReportSummaryType
		avgFine     sql.NullFloat64
		maxFine     sql.NullFloat64
		totalFine   sql.NullFloat64
	)

	summaryQuery := `
		SELECT
			COALESCE(SUM(rt.fine_amount), 0) AS total_fine_amount,
			COUNT(rt.id) AS total_transactions_with_fine,
			COALESCE(AVG(rt.fine_amount), 0) AS average_fine,
			COALESCE(MAX(rt.fine_amount), 0) AS max_fine
		FROM return_transactions rt
		WHERE rt.returned_at >= ? 
		  AND rt.returned_at < ?
		  AND rt.fine_amount > 0
	`

	if err := db.QueryRowContext(
		ctx,
		summaryQuery,
		startDate,
		endDateInclusive,
	).Scan(
		&totalFine,
		&summary.TotalTransactionsWithFine,
		&avgFine,
		&maxFine,
	); err != nil {
		return nil, err
	}

	if totalFine.Valid {
		summary.TotalFineAmount = totalFine.Float64
	}
	if avgFine.Valid {
		summary.AverageFine = avgFine.Float64
	}
	if maxFine.Valid {
		summary.MaxFine = maxFine.Float64
	}

	
	topBorrowers := []types.FineReportTopBorrowerType{}

	topBorrowersQuery := `
		SELECT
			u.id AS borrower_id,
			u.name AS borrower_name,
			COUNT(rt.id) AS fine_count,
			COALESCE(SUM(rt.fine_amount), 0) AS total_fine
		FROM return_transactions rt
		JOIN borrow_transactions bt ON bt.id = rt.borrow_transaction_id
		JOIN users u ON u.id = bt.borrower_id
		WHERE rt.returned_at >= ?
		  AND rt.returned_at < ?
		  AND rt.fine_amount > 0
		GROUP BY u.id, u.name
		ORDER BY total_fine DESC, fine_count DESC
		LIMIT 10
	`

	topRows, err := db.QueryContext(
		ctx,
		topBorrowersQuery,
		startDate,
		endDateInclusive,
	)
	if err != nil {
		return nil, err
	}
	defer topRows.Close()

	for topRows.Next() {
		var item types.FineReportTopBorrowerType

		if err := topRows.Scan(
			&item.BorrowerID,
			&item.BorrowerName,
			&item.FineCount,
			&item.TotalFine,
		); err != nil {
			return nil, err
		}

		topBorrowers = append(topBorrowers, item)
	}

	if err := topRows.Err(); err != nil {
		return nil, err
	}

	
	byDate := []types.FineReportByDateType{}

	byDateQuery := `
		SELECT
			DATE(rt.returned_at) AS fine_date,
			COUNT(rt.id) AS fine_count,
			COALESCE(SUM(rt.fine_amount), 0) AS fine_amount
		FROM return_transactions rt
		WHERE rt.returned_at >= ?
		  AND rt.returned_at < ?
		  AND rt.fine_amount > 0
		GROUP BY DATE(rt.returned_at)
		ORDER BY DATE(rt.returned_at) ASC
	`

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
			item     types.FineReportByDateType
			dateText sql.NullString
		)

		if err := byDateRows.Scan(
			&dateText,
			&item.FineCount,
			&item.FineAmount,
		); err != nil {
			return nil, err
		}

		if dateText.Valid {
			item.Date = dateText.String
		}

		byDate = append(byDate, item)
	}

	if err := byDateRows.Err(); err != nil {
		return nil, err
	}

	return &types.ResGetFineReport{
		Period: types.FineReportPeriodType{
			StartDate: req.StartDate,
			EndDate:   req.EndDate,
		},
		Summary:              summary,
		TopBorrowersWithFine: topBorrowers,
		ByDate:               byDate,
	}, nil
}