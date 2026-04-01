package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/authentication/types"

	"github.com/gin-gonic/gin"
)

func (h *AuthenticationHandler) ChangePasswordHandler(c *gin.Context) {
	log.Printf("[change password] hit sarvice change password with request %v", c.Request)

	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
			"code": http.StatusUnauthorized,
			"error": "missing user id",
		})
		return
	}

	var parsedBody types.DTOChangePassword
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[change password] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	log.Println("[change password] running controller")
	_, err := h.controller.ChangePassword(c.Request.Context(), &types.ReqChangePassword{
		UserID: userID,
		OldPassword: parsedBody.OldPassword,
		NewPassword: parsedBody.NewPassword,
		ConfirmPassword: parsedBody.ConfirmPassword,
	})

	if err != nil {
		msg, code, errMsg := h.utilities.ParseError(err)
		c.JSON(code, gin.H{
			"message": msg,
			"code": code,
			"error": errMsg,
		})
		return
	}

	log.Println("[change password] service change password success")
	c.JSON(http.StatusOK, gin.H{
		"message": "change password successful",
		"code": http.StatusOK,
		"result": nil,
	})
}