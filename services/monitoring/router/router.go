package router

import (
	"services-sipala/services/monitoring/handlers"
	"services-sipala/utilities"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	h *handlers.MonitoringHandler,
	u utilities.IUtility,
) {
	g := r.Group("/monitoring")
}