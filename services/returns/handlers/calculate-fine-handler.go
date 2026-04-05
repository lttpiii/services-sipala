package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/returns/types"

	"github.com/gin-gonic/gin"
)

func (h *ReturnHandler) CalculateFineHandler(c *gin.Context) {
	log.Println("[calculate fine] hit service calculate fine")

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code":    http.StatusForbidden,
			"error":   "required role at least staff",
		})
		return
	}

	var parsedBody types.DTOCalculateFineReturns
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[calculate fine] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	result, err := h.controller.CalculateFine(c.Request.Context(), &types.ReqCalculateFineReturns{
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

	log.Println("[calculate fine] success")
	c.JSON(http.StatusOK, gin.H{
		"message": "Fine calculated",
		"code":    http.StatusOK,
		"result":  result,
	})
}