package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/users/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *UsersHandler) GetListUsersHandler(c *gin.Context) {
	log.Printf("[get list users] hit sarvice get list users with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code": http.StatusForbidden,
			"error": "required role atleast staff",
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	log.Println("[get list users] running controller")
	result, err := h.controller.GetListUsers(c.Request.Context(), &types.ReqGetListUsers{
		Page: page,
		Limit: limit,
		Search: c.Query("search"),
		Role: c.Query("role"),
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

	log.Println("[get list users] service get list users success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get list users successful",
		"code": http.StatusOK,
		"result": result,
	})
}