package handlers

import (
	"services-sipala/services/authentication/controllers"
	"services-sipala/utilities"
)

type AuthenticationHandler struct {
	controller controllers.IAuthenticationController
	utilities utilities.IUtility
}

func New(
	controller controllers.IAuthenticationController,
	utilities utilities.IUtility,
) *AuthenticationHandler {
	return &AuthenticationHandler{
		controller: controller,
		utilities: utilities,
	}
}