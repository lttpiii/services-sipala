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
	g := r.Group("/monitoring", u.AuthMiddleware())

	g.GET("/v1/active-borrows", h.GetActiveBorrowsHandler)
	g.GET("/v1/overdue-borrows", h.GetOverdueBorrowsHandler)
}