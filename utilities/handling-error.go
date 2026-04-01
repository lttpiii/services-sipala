package utilities

import (
	"net/http"
	"strings"
)

func (u *Utility) ParseError(err error) (string, int, string) {
	if err == nil {
		return "success", http.StatusOK, ""
	}

	raw := err.Error()

	// default values
	message := "internal server error"
	code := http.StatusInternalServerError
	detail := raw

	parts := strings.SplitN(raw, ":", 2)

	var key string
	if len(parts) > 0 {
		key = strings.TrimSpace(strings.ToLower(parts[0]))
	}

	if len(parts) == 2 {
		detail = strings.TrimSpace(parts[1])
	}

	// mapping berdasarkan "initial"
	switch key {
	case "invalid request":
		message = "invalid request"
		code = http.StatusBadRequest

	case "not found":
		message = "data not found"
		code = http.StatusNotFound

	case "unauthorized":
		message = "unauthorized"
		code = http.StatusUnauthorized

	case "forbidden":
		message = "forbidden"
		code = http.StatusForbidden

	case "conflict":
		message = "data already exists"
		code = http.StatusConflict
	}

	return message, code, detail
}