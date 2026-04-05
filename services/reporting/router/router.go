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
	g := r.Group("/reports", u.AuthMiddleware())

	 g.GET("/v1/borrows", h.GetBorrowReportHandler)
	 g.GET("/v1/returns", h.GetReturnReportHandler)
	 g.GET("/v1/fines", h.GetFineReportHandler)
}