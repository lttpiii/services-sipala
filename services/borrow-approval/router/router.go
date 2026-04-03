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
	g := r.Group("/approvals", u.AuthMiddleware())

	g.POST("/v1/borrows/:borrow_id/approve", h.ApproveBorrowHandler)
	g.POST("/v1/borrows/:borrow_id/reject", h.RejectBorrowHandler)
}