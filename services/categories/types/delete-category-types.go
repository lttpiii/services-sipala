package types

import "time"

type (
	ReqDeleteCategory struct {
		CategoryID string
	}

	ResDeleteCategory struct {
		ID        string    `json:"id"`
		DeletedAt time.Time `json:"deleted_at"`
	}
)