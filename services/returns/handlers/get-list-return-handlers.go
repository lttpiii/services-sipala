package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/returns/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *ReturnHandler) GetListReturnHandler(c *gin.Context) {
	log.Printf("[get list returns] hit sarvice get list returns with request %v", c.Request)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	log.Println("[get list returns] running controller")
	result, err := h.controller.GetListReturns(c.Request.Context(), &types.ReqGetListReturns{
		Page: page,
		Limit: limit,
		Search: c.Query("search"),
	})

	if err != nil {
		msg, code, errMsg := h.utilities.ParseError(err)
		c.JSON(code, gin.H{
			"message": msg,
			"code": code,
			"error": errMsg,
		})
		return
	}

	log.Println("[get list returns] service get list returns success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get list returns successful",
		"code": http.StatusOK,
		"result": result,
	})
}