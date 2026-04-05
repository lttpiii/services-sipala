package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/returns/types"

	"github.com/gin-gonic/gin"
)

func (h *ReturnHandler) GetReturnByIDHandler(c *gin.Context) {
	log.Printf("[get return by id] hit sarvice get return by id with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code":    http.StatusForbidden,
			"error":   "required role at least staff",
		})
		return
	}

	returnID  := c.Param("id")
	if returnID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing id return",
			"code": http.StatusBadRequest,
			"error": "required id return",
		})
		return
	}

	log.Println("[get return by id] running controller")
	result, err := h.controller.GetReturnByID(c.Request.Context(), &types.ReqGetReturnByID{
		ReturnID: returnID,
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

	log.Println("[get return by id] service get return by id success")
	c.JSON(http.StatusOK, gin.H{
		"message": "get return by id successful",
		"code": http.StatusOK,
		"result": result,
	})
}