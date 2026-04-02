package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/categories/types"

	"github.com/gin-gonic/gin"
)

func (h *CategoriesHandler) DeleteCategoryHandler(c *gin.Context) {
	log.Printf("[delete category] hit sarvice delete category with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code": http.StatusForbidden,
			"error": "required role atleast admin",
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

	log.Println("[delete category] running controller")
	result, err := h.controller.DeleteCategory(c.Request.Context(), &types.ReqDeleteCategory{
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

	log.Println("[delete category] service delete category success")
	c.JSON(http.StatusOK, gin.H{
		"message": "delete category successful",
		"code": http.StatusOK,
		"result": result,
	})
}