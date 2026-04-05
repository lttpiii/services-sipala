package types

type (
	ReqRemoveBorrowItem struct {
		AuthUserRole string
		AuthUserID   string
		BorrowID     string
		ItemID       string
	}

	ResRemoveBorrowItem struct{}
)