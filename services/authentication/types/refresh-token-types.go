package types

type (
	DTORefreshToken struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
)

type (
	ReqRefreshToken struct {
		RefreshToken string
	}

	ResRefreshToken struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
	}
)