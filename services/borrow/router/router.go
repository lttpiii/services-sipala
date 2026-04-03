package router

import (
	"services-sipala/services/borrow/handlers"
	"services-sipala/utilities"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	h *handlers.BorrowHandler,
	u utilities.IUtility,
) {
	g := r.Group("/borrows"),

	g.POST("/v1/borrows"),
	g.GET("/v1/borrows/:id"),
	g.GET("/v1/borrows"),
	g.POST("/v1/borrows/:borrow_id/items/:item_id")	,

}