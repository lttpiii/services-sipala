package utilities

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (u *Utility) GenerateJWT(userID, tokenType, role string, expired time.Duration) (string, string, error) {
	id := uuid.NewString()
	claims := CustomClaims{
		Type: tokenType,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        id,
			Issuer:    "sipala",
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(u.cfg.SecretKey))
	return id, signed, err
}