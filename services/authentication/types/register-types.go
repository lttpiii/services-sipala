package types

type (
	DTORegister struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
)

type (
	ReqRegister struct {
		Name     string
		Email    string
		Password string
	}

	ResRegister struct {
		
	}
)