package types

type (
	ReqGetListCategories struct {
		Page   int
		Limit  int
		Search string
	}

	ResGetListCategories struct {
		Data     []CategoryType `json:"data"`
		Metadata MetadataType   `json:"metadata"`
	}
)