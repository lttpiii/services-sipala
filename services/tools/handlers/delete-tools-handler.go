package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/tools/types"

	"github.com/gin-gonic/gin"
)

func (h *ToolsHandler) DeleteToolHandler(c *gin.Context) {
	id := c.Param("id")
	log.Printf("[delete-tool] hit service delete tool ID: %s\n", id)

	// role := c.GetString("role")
	// if role != "admin" { // Hanya admin sesuai API contract
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "invalid role",
	// 		"code":    http.StatusUnauthorized,
	// 		"error":   "required role admin",
	// 	})
	// 	return
	// }

	log.Println("[delete-tool] running controller")
	result, err := h.controller.DeleteTools(c.Request.Context(), &types.ReqDeleteTool{
		DeleteID: id,
	})

	if err != nil {
		log.Printf("[delete-tool] failed on controller process: %v", err)
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

	log.Printf("[delete-tool] delete tool success")
	c.JSON(http.StatusOK, gin.H{
		"message": "Tool deleted successfully",
		"code":    http.StatusOK,
		"result":  result,
	})
}
