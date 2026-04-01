package reporting

import (
	"services-sipala/config"
	"services-sipala/services/reporting/controllers"
	"services-sipala/services/reporting/handlers"
	"services-sipala/services/reporting/router"
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