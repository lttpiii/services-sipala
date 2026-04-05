package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/tools/types"

	"github.com/gin-gonic/gin"
)

func (h *ToolsHandler) CreateToolHandler(c *gin.Context) {
	log.Printf("[create-tool] hit service create tool with request %v\n", c.Request)

	// role := c.GetString("role")
	// if role != "admin" && role != "staff" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "invalid role",
	// 		"code":    http.StatusUnauthorized,
	// 		"error":   "required role admin or staff",
	// 	})
	// 	return
	// }

	var parsedBody types.DTOCreateTool
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[create-tool] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	log.Println("[create-tool] running controller")
	result, err := h.controller.CreateTools(c.Request.Context(), &types.ReqCreateTool{
		Name:        parsedBody.Name,
		CategoryID:  parsedBody.CategoryID,
		Stock:       parsedBody.Stock,
		Description: parsedBody.Description,
	})

	if err != nil {
		msg, code, errMsg := h.utilities.ParseError(err)

		if err != nil {
			msg, code, errMsg := h.utilities.ParseError(err)
			c.JSON(code, gin.H{
				"message": msg,
				"code":    code,
				"error":   errMsg,
			})
			return
		}

		c.JSON(code, gin.H{
			"message": msg,
			"code":    code,
			"error":   errMsg,
		})
		return
	}

	log.Printf("[create-tool] create tool success")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Tool created successfully",
		"code":    http.StatusCreated,
		"result":  result,
	})
}
