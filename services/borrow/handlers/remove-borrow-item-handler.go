package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/borrow/types"

	"github.com/gin-gonic/gin"
)

func (h *BorrowHandler) RemoveBorrowItemHandler(c *gin.Context) {
	log.Printf("[remove borrow item] hit sarvice remove borrow item with request %v", c.Request)

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

	itemID  := c.Param("item_id")
	if itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id item",
			"code": http.StatusBadRequest,
			"error": "required id item",
		})
		return
	}

	log.Println("[remove borrow item] running controller")
	result, err := h.controller.RemoveBorrowItem(c.Request.Context(), &types.ReqRemoveBorrowItem{
		AuthUserID: authUserID,
		AuthUserRole: authUserRole,
		BorrowID: borrowID,
		ItemID: itemID,
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

	log.Println("[remove borrow item] service remove borrow item success")
	c.JSON(http.StatusOK, gin.H{
		"message": "remove borrow item successful",
		"code": http.StatusOK,
		"result": result,
	})
}