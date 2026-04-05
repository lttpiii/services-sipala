package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/tools/types"

	"github.com/gin-gonic/gin"
)

func (h *ToolsHandler) DeleteToolHandler(c *gin.Context) {
	log.Printf("[delete tool] hit sarvice delete tool with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code": http.StatusForbidden,
			"error": "required role atleast admin",
		})
		return
	}

	authUserID := c.GetString("user_id")

	toolID  := c.Param("id")
	if toolID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id tool",
			"code": http.StatusBadRequest,
			"error": "required id tool",
		})
		return
	}

	log.Println("[delete tool] running controller")
	result, err := h.controller.DeleteTool(c.Request.Context(), &types.ReqDeleteTool{
		AuthUserID: authUserID,
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

	log.Println("[delete tool] service delete tool success")
	c.JSON(http.StatusOK, gin.H{
		"message": "delete tool successful",
		"code": http.StatusOK,
		"result": result,
	})
}