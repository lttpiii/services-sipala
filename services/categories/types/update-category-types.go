package types

type (
	DTOUpdateCategory struct {
		Name string `json:"name" binding:"required"`
	}
)

type (
	ReqUpdateCategory struct {
		AuthUserID string
		CategoryID string
		Name       string
	}

	ResUpdateCategory struct{}
)