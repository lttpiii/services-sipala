package handlers

import (
	"services-sipala/services/borrow/controllers"
	"services-sipala/utilities"
)

type BorrowHandler struct {
	controller controllers.IBorrowController
	utilities utilities.IUtility
}

func New(
	controller controllers.IBorrowController,
	utilities utilities.IUtility,
) *BorrowHandler {
	return &BorrowHandler{
		controller: controller,
		utilities: utilities,
	}
}