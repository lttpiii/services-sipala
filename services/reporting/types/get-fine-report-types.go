package types

type (
	DTOGetFineReport struct {
		StartDate string `form:"start_date" binding:"required"`
		EndDate   string `form:"end_date" binding:"required"`
	}

	ReqGetFineReport struct {
		StartDate string
		EndDate   string
	}

	FinePeriod struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	FineSummary struct {
		TotalFineAmount           float64 `json:"total_fine_amount"`
		TotalTransactionsWithFine int     `json:"total_transactions_with_fine"`
		AverageFine               float64 `json:"average_fine"`
		MaxFine                   float64 `json:"max_fine"`
	}

	FineTopBorrower struct {
		BorrowerID   string  `json:"borrower_id"`
		BorrowerName string  `json:"borrower_name"`
		FineCount    int     `json:"fine_count"`
		TotalFine    float64 `json:"total_fine"`
	}

	FineByDate struct {
		Date       string  `json:"date"`
		FineCount  int     `json:"fine_count"`
		FineAmount float64 `json:"fine_amount"`
	}

	ResGetFineReport struct {
		Period               FinePeriod        `json:"period"`
		Summary              FineSummary       `json:"summary"`
		TopBorrowersWithFine []FineTopBorrower `json:"top_borrowers_with_fine"`
		ByDate               []FineByDate      `json:"by_date"`
	}
)
