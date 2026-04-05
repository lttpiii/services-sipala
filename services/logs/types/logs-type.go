package types

import (
	"time"
)
type(
LogsType struct {
		ID          string        `json:"id"`
		User        *LogUserType  `json:"user,omitempty"`
		Action      string        `json:"action"`
		Entity      string        `json:"entity"`
		EntityID    *string       `json:"entity_id,omitempty"`
		Description *string `json:"description,omitempty"`
		CreatedAt   time.Time     `json:"created_at"`
	}

	LogUserType struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	)