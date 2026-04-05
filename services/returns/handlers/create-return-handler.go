package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/returns/types"

	"github.com/gin-gonic/gin"
)

func (h *ReturnHandler) CreateReturnsHandler(c *gin.Context) {
	log.Printf("[create return] hit service create return with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code":    http.StatusForbidden,
			"error":   "required role at least staff",
		})
		return
	}

	var parsedBody types.DTOCreateReturns
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[create return] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	log.Println("[create return] running controller")
	result, err := h.controller.CreateReturns(c.Request.Context(), &types.ReqCreateReturns{
		BorrowTransactionID: parsedBody.BorrowTransactionID,
	})

	if err != nil {
		msg, code, errMsg := h.utilities.ParseError(err)
		c.JSON(code, gin.H{
			"message": msg,
			"code":    code,
			"error":   errMsg,
		})
		return
	}

	log.Println("[create return] service create return success")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Return transaction created",
		"code":    http.StatusCreated,
		"result":  result,
	})
}