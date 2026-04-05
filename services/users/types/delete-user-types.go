package types

type (
	ReqDeleteUser struct {
		AuthUserID string
		UserID     string
	}

	ResDeleteUser struct{}
)