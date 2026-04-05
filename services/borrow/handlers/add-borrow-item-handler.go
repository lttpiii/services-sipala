package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/borrow/types"

	"github.com/gin-gonic/gin"
)

func (h *BorrowHandler) AddBorrowItemHandler(c *gin.Context) {
	log.Printf("[add borrow item] hit sarvice add borrow item with request %v", c.Request)

	authUserID := c.GetString("user_id")
	authUserRole := c.GetString("role")

	var parsedBody types.DTOAddBorrowItem
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[add borrow item] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	borrowID  := c.Param("borrow_id")
	if borrowID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id borrow",
			"code": http.StatusBadRequest,
			"error": "required id borrow",
		})
		return
	}

	log.Println("[add borrow item] running controller")
	result, err := h.controller.AddBorrowItem(c.Request.Context(), &types.ReqAddBorrowItem{
		AuthUserRole: authUserRole,
		AuthUserID: authUserID,
		BorrowID: borrowID,
		ToolId: parsedBody.ToolId,
		Quantity: parsedBody.Quantity,
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

	log.Println("[add borrow item] service add borrow item success")
	c.JSON(http.StatusCreated, gin.H{
		"message": "add borrow item successful",
		"code": http.StatusCreated,
		"result": result,
	})
}