package main

import (
	"fmt"
	"log"
	"services-sipala/config"
	"services-sipala/services/authentication"
	"services-sipala/services/borrow"
	borrowapproval "services-sipala/services/borrow-approval"
	"services-sipala/services/categories"
	"services-sipala/services/logs"
	"services-sipala/services/tools"
	"services-sipala/services/users"
	"services-sipala/utilities"

	"github.com/gin-gonic/gin"
)

func main() {
	// configuration
	cfg := config.Load()

	// gin
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.UseRawPath = true

	r.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	})

	// grouping base path
	api := r.Group("/api")

	// init utilities
	utils := utilities.New(cfg)

	// init services
	authentication.New(api, cfg, utils)
	users.New(api, cfg, utils)
	categories.New(api, cfg, utils)
	tools.New(api, cfg, utils)
	// borrow.New(api, cfg, utils)
	// borrowapproval.New(api, cfg, utils)
	// returns.New(api, cfg, utils)
	 logs.New(api, cfg, utils)
	// monitoring.New(api, cfg, utils)
	 reporting.New(api, cfg, utils)

	// running server
	log.Printf("server running on port %s", cfg.PORT)
	if err := r.Run(fmt.Sprintf("%s%s", ":", cfg.PORT)); err != nil {
		log.Println("failed to running server")
	}
}