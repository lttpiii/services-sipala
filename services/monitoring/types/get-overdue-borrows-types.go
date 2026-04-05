package types

import "time"

type (
	ReqGetOverdueBorrows struct {
		Page   int    `form:"page"`
		Limit  int    `form:"limit"`
		Search string `form:"search"`
	}

	ResGetOverdueBorrows struct {
		Data     []OverdueBorrowType `json:"data"`
		Metadata MetadataType        `json:"metadata"`
	}

	OverdueBorrowType struct {
		ID            string        `json:"id"`
		Borrower      *BorrowerType `json:"borrower,omitempty"`
		BorrowDate    time.Time     `json:"borrow_date"`
		DueDate       time.Time     `json:"due_date"`
		LateDays      int           `json:"late_days"`
		EstimatedFine int           `json:"estimated_fine"`
		TotalItems    int           `json:"total_items"`
		Items         []ItemsType   `json:"items"`
	}
)