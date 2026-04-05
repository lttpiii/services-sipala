package types

import "time"

type (
	DTOCreateBorrow struct {
		DueDate time.Time `json:"due_date" binding:"required"`
	}
)

type (
	ReqCreateBorrow struct{
		AuthUserID string
		DueDate time.Time	
	}
	
	ResCreateBorrow struct{	
		ID string `json:"id"`
	}
)