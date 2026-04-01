package tools

import (
	"services-sipala/config"
	"services-sipala/services/tools/controllers"
	"services-sipala/services/tools/handlers"
	"services-sipala/services/tools/router"
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