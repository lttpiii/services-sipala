package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/users/types"

	"github.com/gin-gonic/gin"
)

func (h *UsersHandler) CreateUserHandler(c *gin.Context) {
	log.Printf("[create user] hit sarvice create user with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code": http.StatusForbidden,
			"error": "required role atleast staff",
		})
		return
	}


	var parsedBody types.DTOCreateUser
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[create user] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	log.Println("[create user] running controller")
	result, err := h.controller.CreateUser(c.Request.Context(), &types.ReqCreateUser{
		Name:     parsedBody.Name,
		Email:    parsedBody.Email,
		Password: parsedBody.Password,
		Role:     parsedBody.Role,
	})

	if err != nil {
		msg, code, errMsg := h.utilities.ParseError(err)
		c.JSON(code, gin.H{
			"message": msg,
			"code":    code,
			"error":   errMsg,
		})
		return
	}

	log.Println("[create user] service create user success")
	c.JSON(http.StatusCreated, gin.H{
		"message": "create user successful",
		"code":    http.StatusCreated,
		"result":  result,
	})
}
