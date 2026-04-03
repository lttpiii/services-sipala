package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/borrow/types"

	"github.com/gin-gonic/gin"
)

func (h *BorrowHandler) SubmitBorrowHandler(c *gin.Context) {
	log.Printf("[submit borrow] hit sarvice submit borrow with request %v", c.Request)

	borrowID  := c.Param("borrow_id")
	if borrowID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id borrow",
			"code": http.StatusBadRequest,
			"error": "required id borrow",
		})
		return
	}

	log.Println("[submit borrow] running controller")
	result, err := h.controller.SubmitBorrow(c.Request.Context(), &types.ReqSubmitBorrow{
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

	log.Println("[submit borrow] service submit borrow success")
	c.JSON(http.StatusOK, gin.H{
		"message": "submit borrow successful",
		"code": http.StatusOK,
		"result": result,
	})
}