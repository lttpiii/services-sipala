package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/borrow/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *BorrowHandler) GetMyBorrowsHandler(c *gin.Context) {
	log.Printf("[get my borrows] hit sarvice get my borrows with request %v", c.Request)

	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
			"code": http.StatusUnauthorized,
			"error": "missing user id",
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	log.Println("[get my borrows] running controller")
	result, err := h.controller.GetMyBorrows(c.Request.Context(), &types.ReqGetMyBorrows{
		Page: page,
		Limit: limit,
		Status: c.Query("status"),
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

	log.Println("[get my borrows] service get my borrows success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get my borrows successful",
		"code": http.StatusOK,
		"result": result,
	})
}