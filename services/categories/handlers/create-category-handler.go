package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/categories/types"

	"github.com/gin-gonic/gin"
)

func (h *CategoriesHandler) CreateCategoryHandler(c *gin.Context) {
	log.Printf("[create category] hit sarvice create category with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code": http.StatusForbidden,
			"error": "required role atleast staff",
		})
		return
	}

	authUserID := c.GetString("user_id")

	var parsedBody types.DTOCreateCategory
	if err := c.ShouldBindJSON(&parsedBody); err != nil {
		log.Printf("[create category] failed to unmarshal: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"code": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	log.Println("[create category] running controller")
	result, err := h.controller.CreateCategory(c.Request.Context(), &types.ReqCreateCategory{
		AuthUserID: authUserID,
		Name: parsedBody.Name,
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

	log.Println("[create category] service create category success")
	c.JSON(http.StatusCreated, gin.H{
		"message": "create category successful",
		"code": http.StatusCreated,
		"result": result,
	})
}