package types

import "time"

type (
	ReqDeleteTool struct {
		DeleteID string
	}

	ResDeleteUser struct {
		ID        string    `json:"id"`
		DeletedAt time.Time `json:"deleted_at"`
	}
)
