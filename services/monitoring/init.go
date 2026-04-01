package monitoring

import (
	"services-sipala/config"
	"services-sipala/services/monitoring/controllers"
	"services-sipala/services/monitoring/handlers"
	"services-sipala/services/monitoring/router"
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