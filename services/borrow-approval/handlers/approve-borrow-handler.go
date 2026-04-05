package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/borrow-approval/types"

	"github.com/gin-gonic/gin"
)

func (h *BorrowApprovalHandler) ApproveBorrowHandler(c *gin.Context) {
	log.Printf("[approval borrow] hit sarvice approval borrow with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code": http.StatusForbidden,
			"error": "required role atleast staff",
		})
		return
	}

	authUserID := c.GetString("user_id")

	borrowID  := c.Param("borrow_id")
	if borrowID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id borrow",
			"code": http.StatusBadRequest,
			"error": "required id borrow",
		})
		return
	}

	log.Println("[approval borrow] running controller")
	res, err := h.controller.ApproveBorrow(c.Request.Context(), &types.ReqApproveBorrow{
		AuthUserID: authUserID,
		BorrowID: borrowID,
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

	log.Println("[approval borrow] service approval borrow success")
	c.JSON(http.StatusOK, gin.H{
		"message": "approval borrow successful",
		"code": http.StatusOK,
		"result": res,
	})
}