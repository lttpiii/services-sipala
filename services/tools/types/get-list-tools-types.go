package types

type (
	ReqGetListTools struct {
		Page          int
		Limit         int
		Search        string
		CategoryID    string
		AvailableOnly bool
	}

	ResGetListTools struct {
		Data     []ToolType   `json:"data"`
		Metadata MetadataType `json:"metadata"`
	}
)