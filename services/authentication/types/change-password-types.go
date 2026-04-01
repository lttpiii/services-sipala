package types

type (
	DTOChangePassword struct {
		OldPassword     string `json:"old_password" binding:"required,min=6"`
		NewPassword     string `json:"new_password" binding:"required,min=6"`
		ConfirmPassword string `json:"confirm_password" binding:"required,min=6"`
	}
)

type (
	ReqChangePassword struct {
		UserID          string
		OldPassword     string
		NewPassword     string
		ConfirmPassword string
	}

	ResChangePassword struct{}
)