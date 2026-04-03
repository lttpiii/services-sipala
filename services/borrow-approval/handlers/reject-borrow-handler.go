package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/borrow-approval/types"

	"github.com/gin-gonic/gin"
)

func (h *BorrowApprovalHandler) RejectBorrowHandler(c *gin.Context) {
	log.Printf("[reject borrow] hit service reject borrow with request %v", c.Request)

	borrowID := c.Param("borrow_id")
	if borrowID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing borrow_id",
			"code":    http.StatusBadRequest,
			"error":   "required borrow_id",
		})
		return
	}

	var dto types.DTORejectBorrow
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	log.Printf("[reject borrow] processing borrow_id=%s", borrowID)

	res, err := h.controller.RejectBorrow(c.Request.Context(), &types.ReqRejectBorrow{
		BorrowID: borrowID,
		Reason:   dto.Reason,
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

	log.Printf("[reject borrow] success reject borrow", )
	c.JSON(http.StatusOK, gin.H{
		"message": "Borrow transaction rejected",
		"code":    http.StatusOK,
		"result":  res,
	})
}