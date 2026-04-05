package types

type (
	DTOCreateUser struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Role     string `json:"role"`
	}

	ReqCreateUser struct {
		UserID   string
		Name     string
		Email    string
		Password string
		Role     string
	}

	ResCreateUser struct{}
)