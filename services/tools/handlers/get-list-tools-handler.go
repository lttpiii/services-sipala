package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/tools/types"

	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *ToolsHandler) GetListToolsHandler(c *gin.Context) {
	log.Printf("[get list tools] hit sarvice get list tools with request %v", c.Request)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	log.Println("[get list tools] running controller")
	result, err := h.controller.GetListTools(c.Request.Context(), &types.ReqGetListTools{
		Page: page,
		Limit: limit,
		Search: c.Query("search"),
		CategoryID: c.Query("category_id"),
		AvailableOnly: c.Query("available_only") == "true",
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

	log.Println("[get list tools] service get list tools success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get list tools successful",
		"code": http.StatusOK,
		"result": result,
	})
}