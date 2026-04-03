package types

import "time"

type (
	DTOCreateBorrow struct {
		DueDate time.Time `json:"due_date" binding:"required"`
	}
)

type (
	ReqCreateBorrow struct{
		DueDate time.Time	
	}
	
	ResCreateBorrow struct{	
		ID string `json:"id"`
		BorrowID string `json:"borrow_id"`
		Status string `json:"status"`
		BorrowDate string `json:"borrow_date"`
		DueDate time.Time `json:"due_date"`
		Items []ItemsType `json:"items"`
		CreatedAt time.Time `json:"created_at"`
	}
)