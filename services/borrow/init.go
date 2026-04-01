package borrow

import (
	"services-sipala/config"
	"services-sipala/services/borrow/controllers"
	"services-sipala/services/borrow/handlers"
	"services-sipala/services/borrow/router"
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