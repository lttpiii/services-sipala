package types

import "time"

type (
	ToolsType struct {
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		Stok      string     `json:"stock"`
		CreatedAt *time.Time `json:"created_at,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
		DeletedAt *time.Time `json:"deleted_at,omitempty"`
	}
)
