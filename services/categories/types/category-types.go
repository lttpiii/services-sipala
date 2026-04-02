package types

import "time"

type (
	CategoryType struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		ToolsCount *int `json:"tools_count,omitempty"`
		CreatedAt *time.Time `json:"created_at,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}
)