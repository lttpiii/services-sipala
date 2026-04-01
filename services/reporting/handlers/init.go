package handlers

import (
	"services-sipala/services/reporting/controllers"
	"services-sipala/utilities"
)

type ReportingHandler struct {
	controller controllers.IReportingController
	utilities utilities.IUtility
}

func New(
	controller controllers.IReportingController,
	utilities utilities.IUtility,
) *ReportingHandler {
	return &ReportingHandler{
		controller: controller,
		utilities: utilities,
	}
}