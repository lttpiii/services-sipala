package types

type (
	ReqGetOverdueBorrows struct {
		Page   int
		Limit  int
		Search string
	}

	ResGetOverdueBorrows struct {
		Data     []BorrowType `json:"data"`
		Metadata MetadataType `json:"metadata"`
	}
)