package handlers

import (
	"log"
	"net/http"
	"strings"

	"services-sipala/services/reporting/types"

	"github.com/gin-gonic/gin"
)

func (h *ReportingHandler) GetFineReportHandler(c *gin.Context) {
	log.Printf("[get fine report] hit service get fine report with request %v", c.Request)

	role := c.GetString("role")
	if role != "admin" && role != "staff" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid role",
			"code":    http.StatusForbidden,
			"error":   "required role atleast staff",
		})
		return
	}

	startDate := strings.TrimSpace(c.Query("start_date"))
	endDate := strings.TrimSpace(c.Query("end_date"))

	if startDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request query",
			"code":    http.StatusBadRequest,
			"error":   "start_date is required",
		})
		return
	}

	if endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request query",
			"code":    http.StatusBadRequest,
			"error":   "end_date is required",
		})
		return
	}

	log.Println("[get fine report] running controller")
	result, err := h.controller.GetFineReport(c.Request.Context(), &types.ReqGetFineReport{
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		log.Printf("[get fine report] failed on controller process: %v", err)

		msg, code, errMsg := h.utilities.ParseError(err)
		c.JSON(code, gin.H{
			"message": msg,
			"code":    code,
			"error":   errMsg,
		})
		return
	}

	log.Println("[get fine report] service get fine report success")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"code":    http.StatusOK,
		"result":  result,
	})
}