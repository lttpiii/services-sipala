package router

import (
	"services-sipala/services/categories/handlers"
	"services-sipala/utilities"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	h *handlers.CategoriesHandler,
	u utilities.IUtility,
) {
	g := r.Group("/categories")
}