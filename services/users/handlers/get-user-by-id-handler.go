package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/users/types"

	"github.com/gin-gonic/gin"
)

func (h *UsersHandler) GetUserByIDHandler(c *gin.Context) {
	log.Printf("[get user by id] hit sarvice get user by id with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code": http.StatusForbidden,
			"error": "required role atleast staff",
		})
		return
	}

	userID  := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id user",
			"code": http.StatusBadRequest,
			"error": "required id user",
		})
		return
	}

	log.Println("[get user by id] running controller")
	result, err := h.controller.GetUserByID(c.Request.Context(), &types.ReqGetUserByID{
		AuthUserRole: role,
		UserID: userID,
		IncludeDeleted: c.Query("include_deleted") == "true",
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

	log.Println("[get user by id] service get user by id success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get user by id successful",
		"code": http.StatusOK,
		"result": result,
	})
}