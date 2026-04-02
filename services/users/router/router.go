package router

import (
	"services-sipala/services/users/handlers"
	"services-sipala/utilities"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	h *handlers.UsersHandler,
	u utilities.IUtility,
) {
	g := r.Group("/users", u.AuthMiddleware())

	g.POST("/v1/users", h.CreateUserHandler)
	g.PUT("/v1/users/:id", h.UpdateUserHandler)
	g.DELETE("/v1/users/:id", h.DeleteUserHandler)
	g.GET("/v1/users/:id", h.GetUserByIDHandler)
	g.GET("/v1/users", h.GetListUsersHandler)
}