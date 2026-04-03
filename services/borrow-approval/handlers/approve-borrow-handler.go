package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/borrow-approval/types"

	"github.com/gin-gonic/gin"
)

func (h *BorrowApprovalHandler) ApproveBorrowHandler(c *gin.Context) {
	log.Printf("[approval borrow] hit sarvice approval borrow with request %v", c.Request)


	borrowID  := c.Param("id")
	if borrowID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id borrow",
			"code": http.StatusBadRequest,
			"error": "required id borrow",
		})
		return
	}

	log.Println("[approval borrow] running controller")
	_, err := h.controller.ApproveBorrow(c.Request.Context(), &types.ReqApproveBorrow{
		BorrowID: borrowID,
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

	log.Println("[approval borrow] service approval borrow success")
	c.JSON(http.StatusOK, gin.H{
		"message": "approval borrow successful",
		"code": http.StatusOK,
		"result": nil,
	})
}