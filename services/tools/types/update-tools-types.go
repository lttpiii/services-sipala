package types

type (
	DTOUpdateTool struct {
		Name        string `json:"name"`
		CategoryID  string `json:"category_id" binding:"omitempty,uuid"`
		Stock       *int   `json:"stock" binding:"omitempty,min=0"`
		Description string `json:"description"`
	}

	ReqUpdateTool struct {
		ID          string
		Name        string
		CategoryID  string
		Stock       *int
		Description string
	}

	ResUpdateTool ToolsType
)
