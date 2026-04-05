package types

import "time"

type DTOCalculateFineReturns struct {
	BorrowTransactionID string `json:"borrow_transaction_id" binding:"required"`
}

type ReqCalculateFineReturns struct {
	BorrowTransactionID string
}

type ResCalculateFineReturns struct {
	BorrowTransactionID string    `json:"borrow_transaction_id"`
	DueDate             time.Time `json:"due_date"`
	CurrentDate         time.Time `json:"current_date"`
	LateDays            int       `json:"late_days"`
	FinePerDay          float64   `json:"fine_per_day"`
	FineAmount          float64   `json:"fine_amount"`
	Currency            string    `json:"currency"`
}