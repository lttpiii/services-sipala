package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/monitoring/types"

	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *MonitoringHandler) GetOverdueBorrowsHandler(c *gin.Context) {
	log.Printf("[get overdue borrows] hit sarvice get overdue borrows with request %v", c.Request)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	log.Println("[get overdue borrows] running controller")
	result, err := h.controller.GetOverdueBorrows(c.Request.Context(), &types.ReqGetOverdueBorrows{
		Page: page,
		Limit: limit,
		Search: c.Query("search"),
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

	log.Println("[get overdue borrows] service get overdue borrows success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get overdue borrows successful",
		"code": http.StatusOK,
		"result": result,
	})
}