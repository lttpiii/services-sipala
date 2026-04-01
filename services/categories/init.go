package categories

import (
	"services-sipala/config"
	"services-sipala/services/categories/controllers"
	"services-sipala/services/categories/handlers"
	"services-sipala/services/categories/router"
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