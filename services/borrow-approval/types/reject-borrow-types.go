package types

import "time"

type (
	DTORejectBorrow struct {
		Reason string `json:"reason"`
	}
	ReqRejectBorrow struct {
		Reason string
		BorrowID string
	}
	ResRejectBorrow struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		RejectedBy RejectedByType `json:"rejected_by"`
		RejectedAt time.Time `json:"rejected_at"`
		Reason string `json:"reason"`

	}
)

type (
	RejectedByType struct {
		ID   string `json:"rejected_id"`
		Name string `json:"name"`
	}
)