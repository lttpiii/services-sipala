package types

type (
	ReqSubmitBorrow struct {
		AuthUserRole string
		AuthUserID   string
		BorrowID     string
	}

	ResSubmitBorrow struct{}
)