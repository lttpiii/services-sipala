package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/users/types"

	"github.com/gin-gonic/gin"
)

func (h *UsersHandler) DeleteUserHandler(c *gin.Context) {
	log.Printf("[delete user] hit sarvice delete user with request %v", c.Request)

	userID  := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id user",
			"code": http.StatusBadRequest,
			"error": "required id user",
		})
		return
	}

	log.Println("[delete user] running controller")
	result, err := h.controller.DeleteUser(c.Request.Context(), &types.ReqDeleteUser{
		UserID: userID,
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

	log.Println("[delete user] service delete user success")
	c.JSON(http.StatusOK, gin.H{
		"message": "delete user successful",
		"code": http.StatusOK,
		"result": result,
	})
}