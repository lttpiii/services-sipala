package types

type (
	DTOGetReturnReport struct {
		StartDate string `form:"start_date" binding:"required"`
		EndDate   string `form:"end_date" binding:"required"`
	}

	ReqGetReturnReport struct {
		StartDate string
		EndDate   string
	}

	ReturnPeriod struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	ReturnSummary struct {
		TotalReturns int     `json:"total_returns"`
		OnTime       int     `json:"on_time"`
		Late         int     `json:"late"`
		OnTimeRate   float64 `json:"on_time_rate"`
	}

	ReturnByDate struct {
		Date        string `json:"date"`
		ReturnCount int    `json:"return_count"`
		LateCount   int    `json:"late_count"`
	}

	ResGetReturnReport struct {
		Period          ReturnPeriod   `json:"period"`
		Summary         ReturnSummary  `json:"summary"`
		AverageLateDays float64        `json:"average_late_days"`
		ByDate          []ReturnByDate `json:"by_date"`
	}
)
