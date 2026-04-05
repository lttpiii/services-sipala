package types

type (
	DTOAddBorrowItem struct {
		ToolId   string `json:"tool_id" binding:"required"`
		Quantity int    `json:"quantity" binding:"required,min=1"`
	}
)

type (
	ReqAddBorrowItem struct {
		AuthUserID   string
		AuthUserRole string
		BorrowID     string
		ToolId       string
		Quantity     int
	}

	ResAddBorrowItem struct{}
)