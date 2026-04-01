package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/authentication/types"

	"github.com/gin-gonic/gin"
)

func (h *AuthenticationHandler) RegisterHandler(c *gin.Context) {
	log.Printf("[register] hit sarvice register with request %v", c.Request)

	var parsedBody types.DTORegister
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[register] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	log.Println("[register] running controller")
	result, err := h.controller.Register(c.Request.Context(), &types.ReqRegister{
		Name: parsedBody.Name,
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

	log.Println("[register] service register success")
	c.JSON(http.StatusCreated, gin.H{
		"message": "register successful",
		"code": http.StatusCreated,
		"result": result,
	})
}