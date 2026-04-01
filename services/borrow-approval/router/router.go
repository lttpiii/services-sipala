package router

import (
	"services-sipala/services/borrow-approval/handlers"
	"services-sipala/utilities"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	h *handlers.BorrowApprovalHandler,
	u utilities.IUtility,
) {
	g := r.Group("/approvals")
}