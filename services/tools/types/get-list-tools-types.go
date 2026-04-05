package types

type (
	DTOListTools struct {
		Page          int    `form:"page" binding:"omitempty,min=1"`
		Limit         int    `form:"limit" binding:"omitempty,min=1,max=100"`
		Search        string `form:"search"`
		CategoryID    string `form:"category_id" binding:"omitempty,uuid"`
		AvailableOnly bool   `form:"available_only"`
	}

	ReqListTools struct {
		Page          int
		Limit         int
		Search        string
		CategoryID    string
		AvailableOnly bool
	}

	ResListTools struct {
		Data     []ToolsType  `json:"data"`
		Metadata MetadataType `json:"metadata"`
	}
)
