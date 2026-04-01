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
	g := r.Group("/tools")
}