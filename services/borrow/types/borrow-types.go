package types

import "time"

type (
	BorrowType struct {
		ID         string       `json:"id"`
		Borrower   *BorrowerType `json:"borrower,omitempty"`
		Status     string       `json:"status"`
		BorrowDate time.Time `json:"borrow_date"`
		DueDate time.Time `json:"due_date"`
		TotalItems int `json:"total_items"`
		TotalQuantity int `json:"total_quantity"`
		IsOverdue *bool `json:"is_overdue,omitempty"`
	}
)

type (
	BorrowerType struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email *string `json:"email,omitempty"`
	}

	ItemsType struct {
		ID string `json:"id"`
		ToolID string `json:"tool_id"`
		ToolName string `json:"tool_name"`
		Qty int `json:"quantity"`
	}
)