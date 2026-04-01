package handlers

import (
	"services-sipala/services/returns/controllers"
	"services-sipala/utilities"
)

type ReturnHandler struct {
	controller controllers.IReturnController
	utilities utilities.IUtility
}

func New(
	controller controllers.IReturnController,
	utilities utilities.IUtility,
) *ReturnHandler {
	return &ReturnHandler{
		controller: controller,
		utilities: utilities,
	}
}