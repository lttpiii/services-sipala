package types

import "time"

type (
	ReqGetToolByID struct {
		ToolID string
	}

	ResGetToolByID struct {
		ID             string       `json:"id"`
		Name           string       `json:"name"`
		Category       CategoryType `json:"category"`
		Stock          int          `json:"stock"`
		AvailableStock int          `json:"available_stock"`
		BorrowedCount int `json:"borrowed_count"`
		Description    *string      `json:"descriptio"`
		CreatedAt      *time.Time   `json:"created_at"`
		UpdatedAt      *time.Time   `json:"updated_at"`
	}
)