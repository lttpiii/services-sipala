package types

type (
	ReqRemoveBorrowItem struct {
		BorrowID string
		ItemID   string
	}

	ResRemoveBorrowItem struct {
		RemoveitemId string `json:"remove_item_id"`
	}
)