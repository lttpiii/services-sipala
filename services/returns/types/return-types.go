package types

import "time"

type ProcessedBy struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BorrowItem struct {
	ToolName string `json:"tool_name"`
	Quantity int    `json:"quantity"`
}

type BorrowDetails struct {
	BorrowerName string       `json:"borrower_name"`
	BorrowDate   time.Time    `json:"borrow_date"`
	DueDate      time.Time    `json:"due_date"`
	Items        []BorrowItem `json:"items"`
}

type Return struct {
	ID                  string        `json:"id"`
	BorrowTransactionID string        `json:"borrow_transaction_id"`
	ReturnedAt          time.Time     `json:"returned_at"`
	LateDays            *int          `json:"late_days,omitempty"`
	FineAmount          *float64      `json:"fine_amount,omitempty"`
	ProcessedBy         ProcessedBy   `json:"processed_by"`
	BorrowDetails       BorrowDetails `json:"borrow_details"`
	CreatedAt           time.Time     `json:"created_at"`
}