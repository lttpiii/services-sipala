package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/logs/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *LogsHandler) GetLogsHandler(c *gin.Context) {
	log.Printf("[get logs] hit service get logs with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Forbidden",
			"code":    http.StatusForbidden,
			"error":   "required role admin",
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	log.Println("[get logs] running controller")
	result, err := h.controller.GetLogs(c.Request.Context(), &types.ReqGetLogs{
		Page:      page,
		Limit:     limit,
		UserID:    c.Query("user_id"),
		Action:    c.Query("action"),
		Entity:    c.Query("entity"),
		StartDate: c.Query("start_date"),
		EndDate:   c.Query("end_date"),
	})

	if err != nil {
		log.Printf("[get logs] failed on controller process: %v", err)

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
			"code":    code,
			"error":   errMsg,
		})
		return
	}

	log.Println("[get logs] service get logs success")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"code":    http.StatusOK,
		"result":  result,
	})
}