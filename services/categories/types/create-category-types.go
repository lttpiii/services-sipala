package types

type (
	DTOCreateCategory struct {
		Name string `json:"name" binding:"required"`
	}
)

type (
	ReqCreateCategory struct {
		AuthUserID string
		Name       string
	}

	ResCreateCategory struct{}
)