package router

import (
	"services-sipala/services/tools/handlers"
	"services-sipala/utilities"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	h *handlers.ToolsHandler,
	u utilities.IUtility,
) {
	g := r.Group("/tools", u.AuthMiddleware())

	g.POST("/v1/tools", h.CreateToolHandler)
	g.PUT("/v1/tools/:id", h.UpdateToolHandler)
	g.DELETE("/v1/tools/:id", h.DeleteToolHandler)
	g.GET("/v1/tools/:id", h.GetToolByIDHandler)
	g.GET("/v1/tools", h.GetListToolsHandler)
}