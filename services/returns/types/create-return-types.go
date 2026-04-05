package types

import "time"

type (
	DTOCreateReturns struct {
		BorrowTransactionID string `json:"borrow_transaction_id" binding:"required"`
	}
)

type (
	ReqCreateReturns struct {
		AuthUserID          string
		BorrowTransactionID string
	}

	ResCreateReturns struct {
		ID                  string        `json:"id"`
		BorrowTransactionID string        `json:"borrow_transaction_id"`
		ReturnedAt          time.Time     `json:"returned_at"`
		LateDays            int          `json:"late_days"`
		FineAmount          float64      `json:"fine_amount"`
		ProcessedBy         ProcessedBy   `json:"processed_by"`
		BorrowDetails       BorrowDetails `json:"borrow_details"`
		CreatedAt           time.Time     `json:"created_at"`
	}
)
