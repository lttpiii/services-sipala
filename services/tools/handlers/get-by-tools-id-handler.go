package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/tools/types"

	"github.com/gin-gonic/gin"
)

func (h *ToolsHandler) GetToolByIDHandler(c *gin.Context) {
	id := c.Param("id")
	log.Printf("[get-tool-by-id] hit service get tool by id ID: %s\n", id)

	// Auth Bearer (any role), jadi tidak perlu cek spesifik role "admin/staff"

	log.Println("[get-tool-by-id] running controller")
	result, err := h.controller.GetToolsByID(c.Request.Context(), &types.ReqGetToolByID{
		ToolsID: id,
	})

	if err != nil {
		log.Printf("[get-tool-by-id] failed on controller process: %v", err)
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

	log.Printf("[get-tool-by-id] get tool by id success")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"code":    http.StatusOK,
		"result":  result,
	})
}
