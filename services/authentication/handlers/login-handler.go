package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/authentication/types"

	"github.com/gin-gonic/gin"
)

func (h *AuthenticationHandler) LoginHandler(c *gin.Context) {
	log.Printf("[login] hit sarvice login with request %v", c.Request)

	var parsedBody types.DTOLogin
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[login] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	log.Println("[login] running controller")
	result, err := h.controller.Login(c.Request.Context(), &types.ReqLogin{
		Email: parsedBody.Email,
		Password: parsedBody.Password,
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

	log.Println("[login] service login success")
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"code": http.StatusOK,
		"result": result,
	})
}