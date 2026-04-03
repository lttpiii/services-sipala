package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/tools/types"

	"github.com/gin-gonic/gin"
)

func (h *ToolsHandler) UpdateToolHandler(c *gin.Context) {
	id := c.Param("id")
	log.Printf("[update-tool] hit service update tool ID: %s with request %v\n", id, c.Request)

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid role",
			"code":    http.StatusUnauthorized,
			"error":   "required role admin or staff",
		})
		return
	}

	var parsedBody types.DTOUpdateTool
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[update-tool] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	log.Println("[update-tool] running controller")
	result, err := h.controller.UpdateTools(c.Request.Context(), &types.ReqUpdateTool{
		ID:          id,
		Name:        parsedBody.Name,
		CategoryID:  parsedBody.CategoryID,
		Stock:       parsedBody.Stock,
		Description: parsedBody.Description,
	})

	if err != nil {
		log.Printf("[update-tool] failed on controller process: %v", err)

		if err != nil {
			msg, code, errMsg := h.utilities.ParseError(err)
			c.JSON(code, gin.H{
				"message": msg,
				"code":    code,
				"error":   errMsg,
			})
			return
		}

		if err.Error() == "invalid" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid",
				"code":    http.StatusBadRequest,
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"code":    http.StatusInternalServerError,
			"error":   err.Error(),
		})
		return
	}

	log.Printf("[update-tool] update tool success")
	c.JSON(http.StatusOK, gin.H{
		"message": "Tool updated successfully",
		"code":    http.StatusOK,
		"result":  result,
	})
}
