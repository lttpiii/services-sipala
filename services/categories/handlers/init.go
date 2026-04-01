package handlers

import (
	"services-sipala/services/categories/controllers"
	"services-sipala/utilities"
)

type CategoriesHandler struct {
	controller controllers.ICategoriesController
	utilities utilities.IUtility
}

func New(
	controller controllers.ICategoriesController,
	utilities utilities.IUtility,
) *CategoriesHandler {
	return &CategoriesHandler{
		controller: controller,
		utilities: utilities,
	}
}