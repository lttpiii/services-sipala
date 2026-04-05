package handlers

import (
	"log"
	"net/http"
	"services-sipala/services/reporting/types"

	"github.com/gin-gonic/gin"
)

func (h *ReportingHandler) GetBorrowReportHandler(c *gin.Context) {
	log.Printf("[get-borrow-report] hit service get borrow report with query %v\n", c.Request.URL.Query())

	// role := c.GetString("role")
	// if role != "admin" && role != "staff" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "invalid role",
	// 		"code":    http.StatusUnauthorized,
	// 		"error":   "required role admin or staff",
	// 	})
	// 	return
	// }

	var parsedQuery types.DTOGetBorrowReport
	if err := c.ShouldBindQuery(&parsedQuery); err != nil {
		log.Printf("[get-borrow-report] failed to bind query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request query",
			"code":    http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	log.Println("[get-borrow-report] running controller")
	result, err := h.controller.GetBorrowReport(c.Request.Context(), &types.ReqGetBorrowReport{
		StartDate: parsedQuery.StartDate,
		EndDate:   parsedQuery.EndDate,
		GroupBy:   parsedQuery.GroupBy,
	})

	if err != nil {
		log.Printf("[get-borrow-report] failed on controller process: %v", err)

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

	log.Printf("[get-borrow-report] get borrow report success")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"code":    http.StatusOK,
		"result":  result, // Berupa types.ReportType
	})
}
