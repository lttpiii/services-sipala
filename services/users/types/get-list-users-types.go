package types

type (
	ReqGetListUsers struct {
		Page           int
		Limit          int
		Search         string
		Role           string
		IncludeDeleted bool
	}

	ResGetListUsers struct {
		Data     []UserType   `json:"data"`
		Metadata MetadataType `json:"metadata"`
	}
)