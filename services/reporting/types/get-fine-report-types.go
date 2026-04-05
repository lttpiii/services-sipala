package types

type (
	ReqGetFineReport struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	ResGetFineReport struct {
		Period               FineReportPeriodType        `json:"period"`
		Summary              FineReportSummaryType       `json:"summary"`
		TopBorrowersWithFine []FineReportTopBorrowerType `json:"top_borrowers_with_fine"`
		ByDate               []FineReportByDateType      `json:"by_date"`
	}

	FineReportPeriodType struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	FineReportSummaryType struct {
		TotalFineAmount           float64 `json:"total_fine_amount"`
		TotalTransactionsWithFine int     `json:"total_transactions_with_fine"`
		AverageFine               float64 `json:"average_fine"`
		MaxFine                   float64 `json:"max_fine"`
	}

	FineReportTopBorrowerType struct {
		BorrowerID   string  `json:"borrower_id"`
		BorrowerName string  `json:"borrower_name"`
		FineCount    int     `json:"fine_count"`
		TotalFine    float64 `json:"total_fine"`
	}

	FineReportByDateType struct {
		Date       string  `json:"date"`
		FineCount  int     `json:"fine_count"`
		FineAmount float64 `json:"fine_amount"`
	}
)