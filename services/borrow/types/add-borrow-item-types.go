package types

import "time"

type (
	DTOAddBorrowItem struct {
		ToolId string `json:"tool_id" binding:"required"`
		Quantity    int    `json:"quantity" binding:"required,min=1"`
	}
)

type (
	ReqAddBorrowItem struct {

		
		BorrowID string
		ToolId string
		Quantity    int
	}

	ResAddBorrowItem struct {
		ID                  string `json:"id"`
		BorrowTransactionid string `json:"borrow_transaction_id"`
		ToolId              string `json:"tool_id"`
		ToolName            string `json:"tool_name"`
		Quantity                 int    `json:"quantity"`
		CreatedAt           time.Time`json:"created_at"`
	}
)