package utilities

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (u *Utility) UpdateJWT(tokenID, userID, tokenType, role string, expired time.Duration) (string, error) {
	claims := CustomClaims{
		Type: tokenType,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID,
			Issuer:    "sipala",
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(u.cfg.SecretKey))
	return signed, err
}