package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/borrow/types"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *BorrowHandler) GetListBorrowsHandler(c *gin.Context) {
	log.Printf("[get list borrows] hit sarvice get list borrows with request %v", c.Request)

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

	log.Println("[get list borrows] running controller")
	result, err := h.controller.GetListBorrows(c.Request.Context(), &types.ReqGetListBorrows{
		Page: page,
		Limit: limit,
		Status: c.Query("status"),
		BorrowerID: c.Param("borrow_id"),
		StartDate: startDate,
		EndDate: endDate,
		Search: c.Query("search"),
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

	log.Println("[get list borrows] service get list borrows success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get list borrows successful",
		"code": http.StatusOK,
		"result": result,
	})
}