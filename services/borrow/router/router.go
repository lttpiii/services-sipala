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
	g := r.Group("/borrows", u.AuthMiddleware())

	g.POST("/v1/borrows", h.CreateBorrowHandler)
	g.POST("/v1/borrows/:borrow_id/items", h.AddBorrowItemHandler)
	g.DELETE("/v1/borrows/:borrow_id/items/item_id", h.RemoveBorrowItemHandler)
	g.POST("/v1/borrows/:borrow_id/submit", h.SubmitBorrowHandler)
	g.GET("/v1/borrows/:id", h.GetBorrowByIDHandler)
	g.GET("/v1/borrows", h.GetListBorrowsHandler)
	g.GET("/v1/borrows/my-borrows", h.GetMyBorrowsHandler)
}