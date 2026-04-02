package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/users/types"

	"github.com/gin-gonic/gin"
)

func (h *UsersHandler) UpdateUserHandler(c *gin.Context) {
	log.Printf("[update user] hit sarvice update user with request %v", c.Request)

	var parsedBody types.DTOUpdateUser
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[update user] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	userID  := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id user",
			"code": http.StatusBadRequest,
			"error": "required id user",
		})
		return
	}

	log.Println("[update user] running controller")
	result, err := h.controller.UpdateUser(c.Request.Context(), &types.ReqUpdateUser{
		UserID: userID,
		Name: parsedBody.Name,
		Email: parsedBody.Email,
		Role: parsedBody.Role,
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

	log.Println("[update user] service update user success")
	c.JSON(http.StatusOK, gin.H{
		"message": "update user successful",
		"code": http.StatusOK,
		"result": result,
	})
}