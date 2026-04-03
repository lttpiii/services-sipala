package types

import "time"

type (
	ReqGetBorrowByID struct {
		BorrowID string
	}

	ResGetBorrowByID struct {
		ID            string       `json:"id"`
		Borrower      BorrowerType `json:"borrower"`
		Status        string       `json:"status"`
		BorrowDate    time.Time    `json:"borrow_date"`
		DueDate       time.Time    `json:"due_date"`
		Items         []ItemsType `json:"items"`
		ApprovedBy    *string      `json:"approved_by,omitempty"`
		ApprovedAt    *time.Time   `json:"approved_at,omitempty"`
		CreatedAt     time.Time   `json:"created_at"`
	}
)