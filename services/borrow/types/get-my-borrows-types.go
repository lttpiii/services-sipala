package types

type (
	ReqGetMyBorrows struct {
		Page   int
		Limit  int
		Status string
		UserID string
	}

	ResGetMyBorrows struct {
		Data     []BorrowType `json:"data"`
		Metadata MetadataType `json:"metadata"`
	}
)