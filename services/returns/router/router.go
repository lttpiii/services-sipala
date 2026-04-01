package router

import (
	"services-sipala/services/returns/handlers"
	"services-sipala/utilities"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	h *handlers.ReturnHandler,
	u utilities.IUtility,
) {
	g := r.Group("/returns")
}