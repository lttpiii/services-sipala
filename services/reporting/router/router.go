package router

import (
	"services-sipala/services/reporting/handlers"
	"services-sipala/utilities"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	h *handlers.ReportingHandler,
	u utilities.IUtility,
) {
	g := r.Group("/reports")
}