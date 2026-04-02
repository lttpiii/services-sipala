package types

type (
	DTOCreateCategory struct {
		Name string `json:"name" binding:"required"`
	}
)

type (
	ReqCreateCategory struct {
		Name string
	}

	ResCreateCategory CategoryType
)