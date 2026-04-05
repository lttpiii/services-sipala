package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/borrow/types"

	"github.com/gin-gonic/gin"
)

func (h *BorrowHandler) GetBorrowByIDHandler(c *gin.Context) {
	log.Printf("[get borrow by id] hit sarvice get borrow by id with request %v", c.Request)

	authUserID := c.GetString("user_id")
	authUserRole := c.GetString("role")

	borrowID  := c.Param("id")
	if borrowID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id borrow",
			"code": http.StatusBadRequest,
			"error": "required id borrow",
		})
		return
	}

	log.Println("[get borrow by id] running controller")
	result, err := h.controller.GetBorrowByID(c.Request.Context(), &types.ReqGetBorrowByID{
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

	log.Println("[get borrow by id] service get borrow by id success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get borrow by id successful",
		"code": http.StatusOK,
		"result": result,
	})
}