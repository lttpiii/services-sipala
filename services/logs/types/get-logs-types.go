package types

type (
	ReqGetLogs struct {
		Page      int
		Limit     int
		UserID    string
		Action    string
		Entity    string
		StartDate string
		EndDate   string
	}

	ResGetLogs struct {
		Data     []LogsType   `json:"data"`
		Metadata MetadataType `json:"metadata"`
	}
)