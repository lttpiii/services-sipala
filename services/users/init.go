package users

import (
	"services-sipala/config"
	"services-sipala/services/users/controllers"
	"services-sipala/services/users/handlers"
	"services-sipala/services/users/router"
	"services-sipala/utilities"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	cfg *config.Config,
	utilities utilities.IUtility,
) {
	controllers := controllers.New(cfg, utilities)
	handlers := handlers.New(controllers, utilities)
	router.New(r, handlers, utilities)
}