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
	 g.POST("/v1/returns", h.CreateReturnsHandler)
	 g.POST("/v1/returns/calculate-fine", h.CalculateFineHandler)
	 g.GET("/v1/returns/:id", h.GetReturnByIDHandler)
	g.GET("/v1/returns", h.GetListReturnHandler)
}