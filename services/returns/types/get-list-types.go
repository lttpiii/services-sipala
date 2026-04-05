package types

import "time"

type (
	ReqGetListReturns struct {
		Page       int
		Limit      int
		Search     string
		BorrowerID string
		StartDate  *time.Time
		EndDate *time.Time
		HasFine bool
	}

	ResGetListReturns struct {
		Data     []Return     `json:"data"`
		Metadata MetadataType `json:"metadata"`
	}

)