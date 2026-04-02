package types

type (
	ReqGetUserByID struct {
		UserID         string
		IncludeDeleted bool
	}

	ResGetUserByID UserType
)