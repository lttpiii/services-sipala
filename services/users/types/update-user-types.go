package types

type (
	DTOUpdateUser struct {
		Name  string `json:"name"`
		Email string `json:"email" binding:"omitempty,email"`
		Role  string `json:"role"`
	}
)

type (
	ReqUpdateUser struct {
		UserIDOnToken string
		UserRole      string
		UserID        string
		Name          string
		Email         string
		Role          string
	}

	ResUpdateUser struct{}
)