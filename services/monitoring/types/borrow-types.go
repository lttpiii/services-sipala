package types

import "time"

type (
	BorrowType struct {
		ID         string       `json:"id"`
		Borrower   *BorrowerType `json:"borrower,omitempty"`
		Status     string       `json:"status"`
		BorrowDate time.Time `json:"borrow_date"`
		DueDate time.Time `json:"due_date"`
		DaysRemaining int`json:"days_remaining"`
		IsOverdue *bool `json:"is_overdue,omitempty"`
		TotalItems int `json:"total_items"`
		Items []ItemsType
	}
)

type (
	BorrowerType struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email *string `json:"email,omitempty"`
	}

	ItemsType struct {
		ToolName string `json:"tool_name"`
		Qty int `json:"quantity"`
	}
)