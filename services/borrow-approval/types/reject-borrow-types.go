package types

type (
	DTORejectBorrow struct {
		Reason string `json:"reason"`
	}
	ReqRejectBorrow struct {
		AuthUserID string
		Reason     string
		BorrowID   string
	}
	ResRejectBorrow struct{}
)