package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/monitoring/types"

	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *MonitoringHandler) GetActiveBorrowsHandler(c *gin.Context) {
	log.Printf("[get active borrows] hit sarvice get active borrows with request %v", c.Request)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	log.Println("[get active borrows] running controller")
	result, err := h.controller.GetActiveBorrows(c.Request.Context(), &types.ReqGetActiveBorrows{
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

	log.Println("[get active borrows] service get active borrows success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get active borrows successful",
		"code": http.StatusOK,
		"result": result,
	})
}