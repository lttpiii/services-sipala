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
	g := r.Group("/users")
}