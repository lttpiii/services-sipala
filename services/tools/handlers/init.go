package handlers

import (
	"services-sipala/services/tools/controllers"
	"services-sipala/utilities"
)

type ToolsHandler struct {
	controller controllers.IToolsController
	utilities utilities.IUtility
}

func New(
	controller controllers.IToolsController,
	utilities utilities.IUtility,
) *ToolsHandler {
	return &ToolsHandler{
		controller: controller,
		utilities: utilities,
	}
}