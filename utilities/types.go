package utilities

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type (
	CustomClaims struct {
		Type string `json:"type"`
		Role string `json:"role"`
		jwt.RegisteredClaims
	}
)

var (
	ErrTokenExpired   = errors.New("token expired")
	ErrTokenInvalid   = errors.New("token invalid")
	ErrTokenMalformed = errors.New("token malformed")
)

type (
	HTTPError struct {
		Status  int
		Message string
		Code    int
		Error   string
	}
)
