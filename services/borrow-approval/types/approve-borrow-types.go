package types

type (
	ReqApproveBorrow struct {
		AuthUserID string
		BorrowID   string
	}
	ResApproveBorrow struct{}
)