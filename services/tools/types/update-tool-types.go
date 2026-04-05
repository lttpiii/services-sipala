package types

type (
	DTOUpdateTool struct {
		Name        string `json:"name"`
		CategoryID  string `json:"category_id"`
		Stock       int    `json:"stock" binding:"omitempty,min=0"`
		Description string `json:"description"`
	}
)

type (
	ReqUpdateTool struct {
		AuthUserID  string
		ToolID      string
		Name        string
		CategoryID  string
		Stock       int
		Description string
	}

	ResUpdateTool struct{}
)