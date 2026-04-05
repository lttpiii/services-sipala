package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/borrow/types"

	"github.com/gin-gonic/gin"
)

func (h *BorrowHandler) SubmitBorrowHandler(c *gin.Context) {
	log.Printf("[submit borrow] hit sarvice submit borrow with request %v", c.Request)

	authUserID := c.GetString("user_id")
	authUserRole := c.GetString("role")

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
		AuthUserID: authUserID,
		AuthUserRole: authUserRole,
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

	log.Println("[submit borrow] service submit borrow success")
	c.JSON(http.StatusOK, gin.H{
		"message": "submit borrow successful",
		"code": http.StatusOK,
		"result": result,
	})
}