package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/tools/types"

	"github.com/gin-gonic/gin"
)

func (h *ToolsHandler) UpdateToolHandler(c *gin.Context) {
	log.Printf("[update tool] hit sarvice update tool with request %v", c.Request)

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

	var parsedBody types.DTOUpdateTool
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[update tool] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	toolID  := c.Param("id")
	if toolID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id tool",
			"code": http.StatusBadRequest,
			"error": "required id tool",
		})
		return
	}

	log.Println("[update tool] running controller")
	result, err := h.controller.UpdateTool(c.Request.Context(), &types.ReqUpdateTool{
		AuthUserID: authUserID,
		ToolID: toolID,
		Name: parsedBody.Name,
		CategoryID: parsedBody.CategoryID,
		Stock: parsedBody.Stock,
		Description: parsedBody.Description,
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

	log.Println("[update tool] service update tool success")
	c.JSON(http.StatusOK, gin.H{
		"message": "update tool successful",
		"code": http.StatusOK,
		"result": result,
	})
}