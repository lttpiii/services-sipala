package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/authentication/types"

	"github.com/gin-gonic/gin"
)

func (h *AuthenticationHandler) RefreshTokenHandler(c *gin.Context) {
	log.Printf("[refresh token] hit sarvice refresh token with request %v", c.Request)

	var parsedBody types.DTORefreshToken
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[refresh token] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code": http.StatusBadRequest,
			"error": "failed to unmarshal request body: email dan password dibutuhkan, pastikan format email mu benar, dan password minimal 6 karakter",
		})
		return
	}

	log.Println("[refresh token] running controller")
	result, err := h.controller.RefreshToken(c.Request.Context(), &types.ReqRefreshToken{
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

	log.Println("[refresh token] service refresh token success")
	c.JSON(http.StatusOK, gin.H{
		"message": "refresh token successful",
		"code": http.StatusOK,
		"result": result,
	})
}