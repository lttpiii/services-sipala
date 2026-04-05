package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/reporting/types"

	"github.com/gin-gonic/gin"
)

func (h *ReportingHandler) GetFineReportHandler(c *gin.Context) {
	log.Printf("[get-fine-report] hit service get fine report with query %v\n", c.Request.URL.Query())

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid role",
			"code":    http.StatusUnauthorized,
			"error":   "required role admin or staff",
		})
		return
	}

	var parsedQuery types.DTOGetFineReport
	if err := c.ShouldBindQuery(&parsedQuery); err != nil {
		log.Printf("[get-fine-report] failed to bind query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request query",
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	log.Println("[get-fine-report] running controller")
	result, err := h.controller.GetFineReport(c.Request.Context(), &types.ReqGetFineReport{
		StartDate: parsedQuery.StartDate,
		EndDate:   parsedQuery.EndDate,
	})

	if err != nil {
		log.Printf("[get-fine-report] failed on controller process: %v", err)
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

	log.Printf("[get-fine-report] get fine report success")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"code":    http.StatusOK,
		"result":  result, // Berupa types.ReportType
	})
}
