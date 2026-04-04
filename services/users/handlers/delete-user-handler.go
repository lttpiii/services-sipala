package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/users/types"

	"github.com/gin-gonic/gin"
)

func (h *UsersHandler) DeleteUserHandler(c *gin.Context) {
	log.Printf("[delete user] hit sarvice delete user with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code": http.StatusForbidden,
			"error": "required role atleast admin",
		})
		return
	}

	authUserID := c.GetString("user_id")

	userID  := c.Param("id")
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
		AuthUserID: authUserID,
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

	log.Println("[delete user] service delete user success")
	c.JSON(http.StatusOK, gin.H{
		"message": "delete user successful",
		"code": http.StatusOK,
		"result": result,
	})
}