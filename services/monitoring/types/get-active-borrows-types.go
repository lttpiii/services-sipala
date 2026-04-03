package types

type (
	ReqGetActiveBorrows struct {
		Page   int
		Limit  int
		Search string
	}

	ResGetActiveBorrows struct {
		Data     []BorrowType `json:"data"`
		Metadata MetadataType `json:"metadata"`
	}
)