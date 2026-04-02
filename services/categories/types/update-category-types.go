package types

type (
	DTOUpdateCategory struct {
		Name string `json:"name" binding:"required"`
	}
)

type (
	ReqUpdateCategory struct {
		CategoryID string
		Name       string
	}

	ResUpdateCategory CategoryType
)