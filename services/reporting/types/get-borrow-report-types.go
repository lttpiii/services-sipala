package types

type (
	DTOGetBorrowReport struct {
		StartDate string `form:"start_date" binding:"required"`
		EndDate   string `form:"end_date" binding:"required"`
		GroupBy   string `form:"group_by" binding:"omitempty"` // day|week|month
	}

	ReqGetBorrowReport struct {
		StartDate string
		EndDate   string
		GroupBy   string
	}

	BorrowPeriod struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	BorrowSummary struct {
		TotalTransactions  int `json:"total_transactions"`
		TotalItemsBorrowed int `json:"total_items_borrowed"`
		Approved           int `json:"approved"`
		Rejected           int `json:"rejected"`
		Pending            int `json:"pending"`
	}

	BorrowTopTool struct {
		ToolID        string `json:"tool_id"`
		ToolName      string `json:"tool_name"`
		BorrowCount   int    `json:"borrow_count"`
		TotalQuantity int    `json:"total_quantity"`
	}

	BorrowByDate struct {
		Date             string `json:"date"`
		TransactionCount int    `json:"transaction_count"`
		ItemCount        int    `json:"item_count"`
	}

	ResGetBorrowReport struct {
		Period   BorrowPeriod    `json:"period"`
		Summary  BorrowSummary   `json:"summary"`
		TopTools []BorrowTopTool `json:"top_tools"`
		ByDate   []BorrowByDate  `json:"by_date"`
	}
)
