package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/categories/types"

	"github.com/gin-gonic/gin"
)

func (h *CategoriesHandler) UpdateCategoryHandler(c *gin.Context) {
	log.Printf("[update category] hit sarvice update category with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code": http.StatusForbidden,
			"error": "required role atleast staff",
		})
		return
	}

	var parsedBody types.DTOUpdateCategory
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[update category] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	categoryID  := c.Param("id")
	if categoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id category",
			"code": http.StatusBadRequest,
			"error": "required id category",
		})
		return
	}

	log.Println("[update category] running controller")
	result, err := h.controller.UpdateCategory(c.Request.Context(), &types.ReqUpdateCategory{
		CategoryID: categoryID,
		Name: parsedBody.Name,
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

	log.Println("[update category] service update category success")
	c.JSON(http.StatusOK, gin.H{
		"message": "update category successful",
		"code": http.StatusOK,
		"result": result,
	})
}