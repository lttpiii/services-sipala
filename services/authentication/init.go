package authentication

import (
	"services-sipala/config"
	"services-sipala/services/authentication/controllers"
	"services-sipala/services/authentication/handlers"
	"services-sipala/services/authentication/router"
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