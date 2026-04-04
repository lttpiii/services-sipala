package types

type (
	ReqGetUserByID struct {
		AuthUserRole   string
		UserID         string
		IncludeDeleted bool
	}

	ResGetUserByID UserType
)