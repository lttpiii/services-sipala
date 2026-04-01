package types

type (
	DTOLogin struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
)

type (
	ReqLogin struct {
		Email    string
		Password string
	}

	ResLogin struct {
		AccessToken  string   `json:"access_token"`
		RefreshToken string   `json:"refresh_token"`
		TokenType    string   `json:"token_type"`
		ExpiresIn    int      `json:"expires_in"`
		User         UserType `json:"user"`
	}
)