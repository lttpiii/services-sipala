package types

import "time"

type (
	ReqGetListBorrows struct {
		Page       int
		Limit      int
		Status     string
		BorrowerID string
		StartDate  *time.Time
		EndDate *time.Time
		Search string
	}

	ResGetListBorrows struct {
		Data []BorrowType `json:"data"`
		Metadata MetadataType `json:"metadata"`
	}
)