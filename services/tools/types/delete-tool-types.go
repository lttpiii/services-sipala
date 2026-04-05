package types

type (
	ReqDeleteTool struct {
		AuthUserID string
		ToolID     string
	}

	ResDeleteTool struct{}
)