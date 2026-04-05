package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/returns/types"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *ReturnHandler) GetListReturnHandler(c *gin.Context) {
	log.Printf("[get list returns] hit sarvice get list returns with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code":    http.StatusForbidden,
			"error":   "required role at least staff",
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	startDateStr := c.Query("start_date")
	var startDate *time.Time

	if startDateStr != "" {
		parsed, err := time.Parse(time.RFC3339, startDateStr)
		if err != nil {
			log.Printf("failed parse date: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed parse date",
				"code": http.StatusBadRequest,
				"error": err.Error(),
			})
			return
		}
		startDate = &parsed
	}

	endDateStr := c.Query("end_date")
	var endDate *time.Time

	if endDateStr != "" {
		parsed, err := time.Parse(time.RFC3339, endDateStr)
		if err != nil {
			log.Printf("failed parse date: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed parse date",
				"code": http.StatusBadRequest,
				"error": err.Error(),
			})
			return
		}
		endDate = &parsed
	}

	log.Println("[get list returns] running controller")
	result, err := h.controller.GetListReturns(c.Request.Context(), &types.ReqGetListReturns{
		Page: page,
		Limit: limit,
		Search: c.Query("search"),
		BorrowerID: c.Query("borrower_id"),
		StartDate: startDate,
		EndDate: endDate,
		HasFine: c.GetString("has_fine") == "true",
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

	log.Println("[get list returns] service get list returns success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get list returns successful",
		"code": http.StatusOK,
		"result": result,
	})
}