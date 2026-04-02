package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/categories/types"

	"github.com/gin-gonic/gin"
)

func (h *CategoriesHandler) GetCategoryByIDHandler(c *gin.Context) {
	log.Printf("[get category by id] hit sarvice get category by id with request %v", c.Request)

	categoryID  := c.Param("id")
	if categoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id category",
			"code": http.StatusBadRequest,
			"error": "required id category",
		})
		return
	}

	log.Println("[get category by id] running controller")
	result, err := h.controller.GetCategoryByID(c.Request.Context(), &types.ReqGetCategoryByID{
		CategoryID: categoryID,
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

	log.Println("[get category by id] service get category by id success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get category by id successful",
		"code": http.StatusOK,
		"result": result,
	})
}