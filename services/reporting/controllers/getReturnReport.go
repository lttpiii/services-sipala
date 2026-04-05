package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"services-sipala/services/reporting/types"
)

func (c *controller) GetReturnReport(ctx context.Context, req *types.ReqGetReturnReport) (res *types.ResGetReturnReport, err error) {
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
		summary         types.ReturnReportSummaryType
		averageLateDays sql.NullFloat64
	)

	summaryQuery := `
		SELECT
			COUNT(rt.id) AS total_returns,
			COALESCE(SUM(CASE WHEN rt.late_days = 0 THEN 1 ELSE 0 END), 0) AS on_time,
			COALESCE(SUM(CASE WHEN rt.late_days > 0 THEN 1 ELSE 0 END), 0) AS late,
			COALESCE(AVG(CASE WHEN rt.late_days > 0 THEN rt.late_days END), 0) AS average_late_days
		FROM return_transactions rt
		JOIN borrow_transactions bt ON bt.id = rt.borrow_transaction_id
		WHERE rt.returned_at >= ? AND rt.returned_at < ?
	`

	if err := db.QueryRowContext(
		ctx,
		summaryQuery,
		startDate,
		endDateInclusive,
	).Scan(
		&summary.TotalReturns,
		&summary.OnTime,
		&summary.Late,
		&averageLateDays,
	); err != nil {
		return nil, err
	}

	avgLate := 0.0
	if averageLateDays.Valid {
		avgLate = averageLateDays.Float64
	}

	if summary.TotalReturns > 0 {
		summary.OnTimeRate = (float64(summary.OnTime) / float64(summary.TotalReturns)) * 100
	} else {
		summary.OnTimeRate = 0
	}

	byDate := []types.ReturnReportByDateType{}

	byDateQuery := `
		SELECT
			DATE(rt.returned_at) AS return_date,
			COUNT(rt.id) AS return_count,
			COALESCE(SUM(CASE WHEN rt.late_days > 0 THEN 1 ELSE 0 END), 0) AS late_count
		FROM return_transactions rt
		JOIN borrow_transactions bt ON bt.id = rt.borrow_transaction_id
		WHERE rt.returned_at >= ? AND rt.returned_at < ?
		GROUP BY DATE(rt.returned_at)
		ORDER BY DATE(rt.returned_at) ASC
	`

	rows, err := db.QueryContext(
		ctx,
		byDateQuery,
		startDate,
		endDateInclusive,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			item       types.ReturnReportByDateType
			returnDate sql.NullString
		)

		if err := rows.Scan(
			&returnDate,
			&item.ReturnCount,
			&item.LateCount,
		); err != nil {
			return nil, err
		}

		if returnDate.Valid {
			item.Date = returnDate.String
		}

		byDate = append(byDate, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &types.ResGetReturnReport{
		Period: types.ReturnReportPeriodType{
			StartDate: req.StartDate,
			EndDate:   req.EndDate,
		},
		Summary:         summary,
		AverageLateDays: avgLate,
		ByDate:          byDate,
	}, nil
}