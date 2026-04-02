package types

type (
	DTOUpdateUser struct {
		Name  string `json:"name"`
		Email string `json:"email" binding:"email"`
		Role  string `json:"role"`
	}
)

type (
	ReqUpdateUser struct {
		UserID string
		Name   string
		Email  string
		Role   string
	}

	ResUpdateUser UserType
)