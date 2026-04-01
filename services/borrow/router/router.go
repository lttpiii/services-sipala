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
	g := r.Group("/borrows")
}