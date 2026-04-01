package router

import (
	"services-sipala/services/authentication/handlers"
	"services-sipala/utilities"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	h *handlers.AuthenticationHandler,
	u utilities.IUtility,
) {
	g := r.Group("/auth")

	g.POST("/v1/login", h.LoginHandler)
	g.POST("/v1/logout", h.LogoutHandler)
	g.POST("/v1/register", h.RegisterHandler)
	g.POST("/v1/refresh-token", h.RefreshTokenHandler)
	g.GET("/v1/profile", u.AuthMiddleware(), h.GetProfileHandler)
	g.PUT("/v1/change-password", u.AuthMiddleware(), h.ChangePasswordHandler)
}