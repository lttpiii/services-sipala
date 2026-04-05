package types

type (
	ReqGetReturnReport struct {
		StartDate string
		EndDate   string
	}

	ResGetReturnReport struct {
		Period          ReturnReportPeriodType   `json:"period"`
		Summary         ReturnReportSummaryType  `json:"summary"`
		AverageLateDays float64                  `json:"average_late_days"`
		ByDate          []ReturnReportByDateType `json:"by_date"`
	}

	ReturnReportPeriodType struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	ReturnReportSummaryType struct {
		TotalReturns int     `json:"total_returns"`
		OnTime       int     `json:"on_time"`
		Late         int     `json:"late"`
		OnTimeRate   float64 `json:"on_time_rate"`
	}

	ReturnReportByDateType struct {
		Date        string `json:"date"`
		ReturnCount int    `json:"return_count"`
		LateCount   int    `json:"late_count"`
	}
)