package controllers

import (
	"context"
	"fmt"
	"math"
	"services-sipala/services/users/types"
	"strings"
)

func (c *controller) GetListUsers(ctx context.Context, req *types.ReqGetListUsers) (res *types.ResGetListUsers, err error) {
	db := c.cfg.MysqlClient

	var (
		users []types.UserType

		conds []string
		args []any
		order string
		totalRecords int
	)

	order = "ORDER BY created_at DESC"
	
	if !req.IncludeDeleted {
		conds = append(conds, "deleted_at is null")
	}

	if req.Search != "" {
		conds = append(conds, "name like ?")
		args = append(args, "%"+req.Search+"%")
	}

	if req.Role != "" {
		conds = append(conds, "role = ?")
		args = append(args, req.Role)
	}

	whereClause := ""
	if len(conds) > 0 {
		whereClause = "WHERE " + strings.Join(conds, " AND ")
	}

	countArgs := append([]any{}, args...)

	if err := db.QueryRowContext(
		ctx,
		fmt.Sprintf("select count(*) from users %s", whereClause),
		countArgs...
	).Scan(&totalRecords); err != nil {
		return nil, err
	}

	offset := (req.Page - 1) * req.Limit
	dataArgs := append(args, req.Limit, offset)

	dataQuery := fmt.Sprintf(`
	select
		id,
		name,
		email,
		role,
		created_at
	from users
	%s
	%s
	limit ? offset ?
	`, whereClause, order)

	rows, err := db.QueryContext(
		ctx, dataQuery, dataArgs...
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u types.UserType

		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Role,
			&u.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	totalPages := int(math.Ceil(float64(totalRecords)/ float64(req.Limit)))

	return &types.ResGetListUsers{
		Data: users,
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