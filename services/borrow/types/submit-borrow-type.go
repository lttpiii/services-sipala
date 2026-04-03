package types

import "time"

type (
	ReqSubmitBorrow struct{
		BorrowID string
	}

	ResSubmitBorrow struct {
		ID          string `json:"id"`
		Status      string `json:"status"`
		SubmittedAt time.Time`json:"submitted_at"`
		TotalItems int `json:"total_items"`
		TotalQty int `json:"total_qty"`
	}
)