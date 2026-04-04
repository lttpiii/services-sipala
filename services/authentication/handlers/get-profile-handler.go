package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/authentication/types"

	"github.com/gin-gonic/gin"
)

func (h *AuthenticationHandler) GetProfileHandler(c *gin.Context) {
	log.Printf("[get profile] hit sarvice get profile with request %v", c.Request)

	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
			"code": http.StatusUnauthorized,
			"error": "missing user id",
		})
		return
	}

	log.Println("[get profile] running controller")
	result, err := h.controller.GetProfile(c.Request.Context(), &types.ReqGetProfile{
		UserID: userID,
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

	log.Println("[get profile] service get profile success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get profile successful",
		"code": http.StatusOK,
		"result": result,
	})
}