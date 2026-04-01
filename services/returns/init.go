package returns

import (
	"services-sipala/config"
	"services-sipala/services/returns/controllers"
	"services-sipala/services/returns/handlers"
	"services-sipala/services/returns/router"
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
