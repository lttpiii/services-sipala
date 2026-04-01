package handlers

import (
	"services-sipala/services/logs/controllers"
	"services-sipala/utilities"
)

type LogsHandler struct {
	controller controllers.ILogsController
	utilities utilities.IUtility
}

func New(
	controller controllers.ILogsController,
	utilities utilities.IUtility,
) *LogsHandler {
	return &LogsHandler{
		controller: controller,
		utilities: utilities,
	}
}