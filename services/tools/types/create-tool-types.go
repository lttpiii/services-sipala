package types

type (
	DTOCreateTool struct {
		Name        string `json:"name" binding:"required"`
		CategoryID  string `json:"category_id" binding:"required"`
		Stock       int    `json:"stock" binding:"required,min=0"`
		Description string `json:"description"`
	}
)

type (
	ReqCreateTool struct {
		AuthUserID  string
		Name        string
		CategoryID  string
		Stock       int
		Description string
	}

	ResCreateTool struct{}
)