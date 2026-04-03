package types

import "time"

type (
	DTOCreateTool struct {
		Name        string `json:"name" binding:"required"`
		CategoryID  string `json:"category_id" binding:"required,uuid"`
		Stock       int    `json:"stock" binding:"required,min=0"`
		Description string `json:"description"`
	}

	ReqCreateTool struct {
		Name        string
		CategoryID  string
		Stock       int
		Description string
	}

	ResCreateTool struct {
		ID           string    `json:"id"`
		Name         string    `json:"name"`
		CategoryID   string    `json:"category_id"`
		CategoryName string    `json:"category_name"`
		Stock        int       `json:"stock"`
		Description  string    `json:"description"`
		CreatedAt    time.Time `json:"created_at"`
	}
)
