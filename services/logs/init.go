package logs

import (
	"services-sipala/config"
	"services-sipala/services/logs/controllers"
	"services-sipala/services/logs/handlers"
	"services-sipala/services/logs/router"
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