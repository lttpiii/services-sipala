package utilities

import (
	"time"

	"github.com/gin-gonic/gin"
)

type (
	IUtility interface {
		// JWT Token
		GenerateJWT(userID, tokenType, role string, expired time.Duration) (string, string, error)
		ValidateJWT(tokenString string) (*CustomClaims, error)
		UpdateJWT(tokenID, userID, tokenType, role string, expired time.Duration) (string, error)

		// Error Handling
		ParseError(err error) (string, int, string)

		// Hash
		CompareStringWithHash(plain, hash string) bool
		HashString(plain string) (string, error)

		// Middleware
		AuthMiddleware() gin.HandlerFunc
	}
)