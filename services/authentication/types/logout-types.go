package types

type (
	DTOLogout struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
)

type (
	ReqLogout struct {
		RefreshToken string
	}

	ResLogout UserType
)