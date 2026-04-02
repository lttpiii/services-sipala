package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/categories/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *CategoriesHandler) GetListCategoriesHandler(c *gin.Context) {
	log.Printf("[get list categories] hit sarvice get list categories with request %v", c.Request)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	log.Println("[get list categories] running controller")
	result, err := h.controller.GetListCategories(c.Request.Context(), &types.ReqGetListCategories{
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

	log.Println("[get list categories] service get list categories success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get list categories successful",
		"code": http.StatusOK,
		"result": result,
	})
}