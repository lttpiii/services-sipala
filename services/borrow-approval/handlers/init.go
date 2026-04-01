package handlers

import (
	"services-sipala/services/borrow-approval/controllers"
	"services-sipala/utilities"
)

type BorrowApprovalHandler struct {
	controller controllers.IBorrowApprovalController
	utilities utilities.IUtility
}

func New(
	controller controllers.IBorrowApprovalController,
	utilities utilities.IUtility,
) *BorrowApprovalHandler {
	return &BorrowApprovalHandler{
		controller: controller,
		utilities: utilities,
	}
}