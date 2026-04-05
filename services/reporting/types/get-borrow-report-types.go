package types

type (
	ReqGetBorrowReport struct {
		StartDate string
		EndDate   string
		GroupBy   string
	}

	ResGetBorrowReport struct {
		Period   BorrowReportPeriodType    `json:"period"`
		Summary  BorrowReportSummaryType   `json:"summary"`
		TopTools []BorrowReportTopToolType `json:"top_tools"`
		ByDate   []BorrowReportByDateType  `json:"by_date"`
	}

	BorrowReportPeriodType struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	BorrowReportSummaryType struct {
		TotalTransactions  int `json:"total_transactions"`
		TotalItemsBorrowed int `json:"total_items_borrowed"`
		Approved           int `json:"approved"`
		Rejected           int `json:"rejected"`
		Pending            int `json:"pending"`
	}

	BorrowReportTopToolType struct {
		ToolID        string `json:"tool_id"`
		ToolName      string `json:"tool_name"`
		BorrowCount   int    `json:"borrow_count"`
		TotalQuantity int    `json:"total_quantity"`
	}

	BorrowReportByDateType struct {
		Date             string `json:"date"`
		TransactionCount int    `json:"transaction_count"`
		ItemCount        int    `json:"item_count"`
	}
)