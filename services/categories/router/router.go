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
	g := r.Group("/categories", u.AuthMiddleware())

	g.POST("/v1/categories", h.CreateCategoryHandler)
	g.PUT("/v1/categories/:id", h.UpdateCategoryHandler)
	g.DELETE("/v1/categories/:id", h.DeleteCategoryHandler)
	g.GET("/v1/categories/:id", h.GetCategoryByIDHandler)
	g.GET("/v1/categories", h.GetListCategoriesHandler)
}