package utilities

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (u *Utility) ValidateJWT(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrTokenInvalid
			}
			return []byte(u.cfg.SecretKey), nil
		},
		jwt.WithIssuer("sipala"),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, ErrTokenExpired
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, ErrTokenMalformed
		default:
			return nil, ErrTokenInvalid
		}
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, ErrTokenInvalid
	}

	// VALIDASI MANUAL TAMBAHAN (tanpa cek Type)
	if claims.Subject == "" || claims.Role == "" {
		return nil, ErrTokenInvalid
	}

	if claims.IssuedAt == nil || claims.ExpiresAt == nil {
		return nil, ErrTokenInvalid
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, ErrTokenExpired
	}

	return claims, nil
}