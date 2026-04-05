package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"services-sipala/services/logs/types"
	"strings"
	"time"
)

func (c *controller) GetLogs(ctx context.Context, req *types.ReqGetLogs) (res *types.ResGetLogs, err error) {
	db := c.cfg.MysqlClient

var (
	conds        []string
	args         []any
	order        string
	totalRecords int
)

logs := make([]types.LogsType, 0)

	if req.Page < 1 {
		req.Page = 1
	}

	if req.Limit < 1 {
		req.Limit = 10
	}

	if req.Limit > 100 {
		req.Limit = 100
	}

	if req.Action != "" {
		req.Action = strings.ToUpper(req.Action)

		validActions := map[string]bool{
			"CREATE": true,
			"UPDATE": true,
			"DELETE": true,
			"LOGIN":  true,
			"LOGOUT": true,
		}

		if !validActions[req.Action] {
			return nil, fmt.Errorf("invalid action filter")
		}
	}

	if req.Entity != "" {
		validEntities := map[string]bool{
			"users":               true,
			"tools":               true,
			"categories":          true,
			"borrow_transactions": true,
			"return_transactions": true,
		}

		if !validEntities[req.Entity] {
			return nil, fmt.Errorf("invalid entity filter")
		}
	}

	if req.StartDate != "" {
		if _, err := time.Parse("2006-01-02", req.StartDate); err != nil {
			return nil, fmt.Errorf("invalid start_date format, use YYYY-MM-DD")
		}
	}

	if req.EndDate != "" {
		if _, err := time.Parse("2006-01-02", req.EndDate); err != nil {
			return nil, fmt.Errorf("invalid end_date format, use YYYY-MM-DD")
		}
	}

	if req.StartDate != "" && req.EndDate != "" {
		startDate, _ := time.Parse("2006-01-02", req.StartDate)
		endDate, _ := time.Parse("2006-01-02", req.EndDate)

		if endDate.Before(startDate) {
			return nil, fmt.Errorf("end_date cannot be before start_date")
		}
	}

	order = "ORDER BY al.created_at DESC"

	if req.UserID != "" {
		conds = append(conds, "al.user_id = ?")
		args = append(args, req.UserID)
	}

	if req.Action != "" {
		conds = append(conds, "al.action = ?")
		args = append(args, req.Action)
	}

	if req.Entity != "" {
		conds = append(conds, "al.entity = ?")
		args = append(args, req.Entity)
	}

	if req.StartDate != "" {
	conds = append(conds, "al.created_at >= ?")
	args = append(args, req.StartDate+" 00:00:00")
}

if req.EndDate != "" {
	conds = append(conds, "al.created_at <= ?")
	args = append(args, req.EndDate+" 23:59:59")
}

	whereClause := ""
	if len(conds) > 0 {
		whereClause = "WHERE " + strings.Join(conds, " AND ")
	}

	countArgs := append([]any{}, args...)

	countQuery := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM activity_logs al
		LEFT JOIN users u ON u.id = al.user_id
		%s
	`, whereClause)

	if err := db.QueryRowContext(
		ctx,
		countQuery,
		countArgs...,
	).Scan(&totalRecords); err != nil {
		return nil, err
	}

	offset := (req.Page - 1) * req.Limit
	dataArgs := append(args, req.Limit, offset)

	dataQuery := fmt.Sprintf(`
		SELECT
			al.id,
			u.id,
			u.name,
			al.action,
			al.entity,
			al.entity_id,
			al.description,
			al.created_at
		FROM activity_logs al
		LEFT JOIN users u ON u.id = al.user_id
		%s
		%s
		LIMIT ? OFFSET ?
	`, whereClause, order)

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
		var logItem types.LogsType
		var userID sql.NullString
		var userName sql.NullString
		var entityID sql.NullString
		var description sql.NullString

		if err := rows.Scan(
			&logItem.ID,
			&userID,
			&userName,
			&logItem.Action,
			&logItem.Entity,
			&entityID,
			&description,
			&logItem.CreatedAt,
		); err != nil {
			return nil, err
		}

		if userID.Valid && userName.Valid {
	logItem.User = &types.LogUserType{
		ID:   userID.String,
		Name: userName.String,
	}
}

		if entityID.Valid {
			logItem.EntityID = &entityID.String
		}

		if description.Valid {
	logItem.Description = &description.String
}

		logs = append(logs, logItem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	totalPages := 0
	if totalRecords > 0 {
		totalPages = int(math.Ceil(float64(totalRecords) / float64(req.Limit)))
	}

	return &types.ResGetLogs{
		Data: logs,
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