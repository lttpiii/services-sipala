package types

type (
	ReqDeleteCategory struct {
		AuthUserID string
		CategoryID string
	}

	ResDeleteCategory struct{}
)