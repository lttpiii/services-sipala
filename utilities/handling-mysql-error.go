package utilities

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func (u *Utility) ParseMySQLError(err error) *HTTPError {
	var mysqlErr *mysql.MySQLError

	// 🔥 penting banget
	if !errors.As(err, &mysqlErr) {
		return nil
	}

	switch mysqlErr.Number {

	// =====================
	// VALIDATION ERRORS
	// =====================

	case 1048: // Column cannot be null
		return &HTTPError{
			Status:  http.StatusBadRequest,
			Code:    http.StatusBadRequest,
			Message: "missing required field",
			Error:   "not null constraint violated",
		}

	case 3819: // Check constraint failed (MySQL 8+)
		return &HTTPError{
			Status:  http.StatusBadRequest,
			Code:    http.StatusBadRequest,
			Message: "validation failed",
			Error:   "check constraint violated",
		}

	// =====================
	// CONFLICT ERRORS
	// =====================

	case 1062: // Duplicate entry
		msg := "resource already exists"

		// 🔥 powerful: parse index name dari message
		if strings.Contains(mysqlErr.Message, "users.email") {
			msg = "email already registered"
		} else if strings.Contains(mysqlErr.Message, "users.username") {
			msg = "username already taken"
		}

		return &HTTPError{
			Status:  http.StatusConflict,
			Code:    http.StatusConflict,
			Message: msg,
			Error:   "duplicate entry",
		}

	case 1452: // Cannot add/update child row (FK)
		return &HTTPError{
			Status:  http.StatusConflict,
			Code:    http.StatusConflict,
			Message: "invalid reference",
			Error:   "foreign key constraint violated",
		}

	case 1451: // Cannot delete/update parent row
		return &HTTPError{
			Status:  http.StatusConflict,
			Code:    http.StatusConflict,
			Message: "resource is still in use",
			Error:   "foreign key constraint violated",
		}

	// =====================
	// AUTH & PERMISSION
	// =====================

	case 1044, 1045: // access denied
		return &HTTPError{
			Status:  http.StatusForbidden,
			Code:    http.StatusForbidden,
			Message: "access denied",
			Error:   "insufficient privilege",
		}

	// =====================
	// TRANSACTION
	// =====================

	case 1213, // deadlock
		1205: // lock wait timeout
		return &HTTPError{
			Status:  http.StatusServiceUnavailable,
			Code:    http.StatusServiceUnavailable,
			Message: "please retry request",
			Error:   "transaction conflict",
		}

	// =====================
	// TIMEOUT
	// =====================

	case 1317: // query execution interrupted
		return &HTTPError{
			Status:  http.StatusRequestTimeout,
			Code:    http.StatusRequestTimeout,
			Message: "request timeout",
			Error:   "query interrupted",
		}

	// =====================
	// SCHEMA / QUERY ERROR
	// =====================

	case 1146, // table doesn't exist
		1054, // unknown column
		1064: // syntax error
		return &HTTPError{
			Status:  http.StatusInternalServerError,
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
			Error:   "database schema error",
		}

	// =====================
	// CONNECTION
	// =====================

	case 2002, 2006, 2013: // connection errors
		return &HTTPError{
			Status:  http.StatusServiceUnavailable,
			Code:    http.StatusServiceUnavailable,
			Message: "database unavailable",
			Error:   "connection failure",
		}
	}

	// fallback mysql error
	return &HTTPError{
		Status:  http.StatusInternalServerError,
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
		Error:   "mysql error",
	}
}