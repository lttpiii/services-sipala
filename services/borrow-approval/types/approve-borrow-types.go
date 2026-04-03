package types

import "time"

type (
	ReqApproveBorrow struct {
		BorrowID string
	}
	ResApproveBorrow struct {
		ID         string         `json:"id"`
		Status     string         `json:"status"`
		Approvedby ApprovedbyType `json:"approved_by"`
		ApprovedAt time.Time `json:"approved_at"`
		BorrowDate time.Time `json:"borrow_date"`
		DueDate time.Time `json:"due_date"`
	}
)
type (
	ApprovedbyType struct {
		ID   string `json:"id"`
		Name string `json:"nama"`
	}
)