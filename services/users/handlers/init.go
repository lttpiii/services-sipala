package handlers

import (
	"services-sipala/services/users/controllers"
	"services-sipala/utilities"
)

type UsersHandler struct {
	controller controllers.IUsersController
	utilities utilities.IUtility
}

func New(
	controller controllers.IUsersController,
	utilities utilities.IUtility,
) *UsersHandler {
	return &UsersHandler{
		controller: controller,
		utilities: utilities,
	}
}