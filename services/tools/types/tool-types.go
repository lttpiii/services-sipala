package types

type (
	ToolType struct {
		ID             string       `json:"id"`
		Name           string       `json:"name"`
		Category       CategoryType `json:"category"`
		Stock          int          `json:"stock"`
		AvailableStock int          `json:"available_stock"`
		Description    *string      `json:"description"`
	}
)

type (
	CategoryType struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)