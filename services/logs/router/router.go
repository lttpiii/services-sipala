package router

import (
	"services-sipala/services/logs/handlers"
	"services-sipala/utilities"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	h *handlers.LogsHandler,
	u utilities.IUtility,
) {
	g := r.Group("/logs",u.AuthMiddleware())
	
	g.GET("/v1/logs", h.GetLogsHandler)
}