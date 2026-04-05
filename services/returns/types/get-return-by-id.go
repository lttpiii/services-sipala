package types

import "time"

type (
	ReqGetReturnByID struct {
		ReturnID string
	}

	ResGetReturnByID struct {
		ID                  string        `json:"id"`
		BorrowTransactionID string        `json:"borrow_transaction_id"`
		ReturnedAt          time.Time     `json:"returned_at"`
		LateDays            *int          `json:"late_days,omitempty"`
		FineAmount          *float64      `json:"fine_amount,omitempty"`
		ProcessedBy         ProcessedBy   `json:"processed_by"`
		BorrowDetails       BorrowDetails `json:"borrow_details"`
		CreatedAt           time.Time     `json:"created_at"`
	}
)