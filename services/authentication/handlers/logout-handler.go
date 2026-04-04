package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/authentication/types"

	"github.com/gin-gonic/gin"
)

func (h *AuthenticationHandler) LogoutHandler(c *gin.Context) {
	log.Printf("[logout] hit sarvice logout with request %v", c.Request)

	var parsedBody types.DTOLogout
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[logout] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	log.Println("[logout] running controller")
	_, err := h.controller.Logout(c.Request.Context(), &types.ReqLogout{
		RefreshToken: parsedBody.RefreshToken,
	})

	if err != nil {
		log.Printf("failed on controller process: %v", err)

		if mysqlError := h.utilities.ParseMySQLError(err); mysqlError != nil {
			c.JSON(mysqlError.Status, gin.H{
				"message": mysqlError.Message,
				"code":    mysqlError.Code,
				"error":   mysqlError.Error,
			})
			return
		}

		msg, code, errMsg := h.utilities.ParseError(err)
		c.JSON(code, gin.H{
			"message": msg,
			"code": code,
			"error": errMsg,
		})
		return
	}

	log.Println("[logout] service logout success")
	c.JSON(http.StatusOK, gin.H{
		"message": "logout successful",
		"code": http.StatusOK,
		"result": nil,
	})
}