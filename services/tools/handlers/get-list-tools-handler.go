package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/tools/types"

	"github.com/gin-gonic/gin"
)

func (h *ToolsHandler) ListToolsHandler(c *gin.Context) {
	log.Printf("[list-tools] hit service list tools with query %v\n", c.Request.URL.Query())

	// Auth Bearer (any role), tidak perlu cek role spesifik

	var parsedQuery types.DTOListTools
	if err := c.ShouldBindQuery(&parsedQuery); err != nil {
		log.Printf("[list-tools] failed to unmarshal query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request query",
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	log.Println("[list-tools] running controller")
	result, err := h.controller.ListTools(c.Request.Context(), &types.ReqListTools{
		Page:          parsedQuery.Page,
		Limit:         parsedQuery.Limit,
		Search:        parsedQuery.Search,
		CategoryID:    parsedQuery.CategoryID,
		AvailableOnly: parsedQuery.AvailableOnly,
	})

	if err != nil {
		log.Printf("[list-tools] failed on controller process: %v", err)
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

	log.Printf("[list-tools] list tools success")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"code":    http.StatusOK,
		"result":  result,
	})
}
