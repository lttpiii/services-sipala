package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/tools/types"

	"github.com/gin-gonic/gin"
)

func (h *ToolsHandler) GetToolByIDHandler(c *gin.Context) {
	log.Printf("[get tool by id] hit sarvice get tool by id with request %v", c.Request)

	toolID  := c.Param("id")
	if toolID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id tool",
			"code": http.StatusBadRequest,
			"error": "required id tool",
		})
		return
	}

	log.Println("[get tool by id] running controller")
	result, err := h.controller.GetToolByID(c.Request.Context(), &types.ReqGetToolByID{
		ToolID: toolID,
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

	log.Println("[get tool by id] service get tool by id success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get tool by id successful",
		"code": http.StatusOK,
		"result": result,
	})
}