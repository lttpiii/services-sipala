package types

import "time"

type (
	ReqDeleteUser struct {
		UserID string
	}

	ResDeleteUser struct {
		ID        string `json:"id"`
		DeletedAt time.Time `json:"deleted_at"`
	}
)