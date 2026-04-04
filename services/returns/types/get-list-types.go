package types

type (
	ReqGetListReturns struct {
		Page   int
		Limit  int
		Search string
	}

	ResGetListReturns struct {
		Data     []Return     `json:"data"`
		Metadata MetadataType `json:"metadata"`
	}
)