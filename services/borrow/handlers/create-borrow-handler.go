package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/borrow/types"

	"github.com/gin-gonic/gin"
)

func (h *BorrowHandler) CreateBorrowHandler(c *gin.Context) {
	log.Printf("[create borrow] hit sarvice create borrow with request %v", c.Request)

	var parsedBody types.DTOCreateBorrow
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[create borrow] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	log.Println("[create borrow] running controller")
	result, err := h.controller.CreateBorrow(c.Request.Context(), &types.ReqCreateBorrow{
		DueDate: parsedBody.DueDate,
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

	log.Println("[create borrow] service create borrow success")
	c.JSON(http.StatusCreated, gin.H{
		"message": "create borrow successful",
		"code": http.StatusCreated,
		"result": result,
	})
}