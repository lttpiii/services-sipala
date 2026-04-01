package handlers

import (
	"services-sipala/services/monitoring/controllers"
	"services-sipala/utilities"
)

type MonitoringHandler struct {
	controller controllers.IMonitoringController
	utilities utilities.IUtility
}

func New(
	controller controllers.IMonitoringController,
	utilities utilities.IUtility,
) *MonitoringHandler {
	return &MonitoringHandler{
		controller: controller,
		utilities: utilities,
	}
}