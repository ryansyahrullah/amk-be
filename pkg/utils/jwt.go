package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims mendeskripsikan payload token.
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	NRP      string `json:"nrp"`
	RoleSlug string `json:"role_slug"`
	jwt.RegisteredClaims
}

// GenerateToken membuat token JWT sederhana.
func GenerateToken(secret string, userID uint, nrp, roleSlug string, ttl time.Duration) (string, error) {
	if secret == "" {
		return "", errors.New("jwt secret is empty")
	}

	claims := &JWTClaims{
		UserID:   userID,
		NRP:      nrp,
		RoleSlug: roleSlug,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken membaca token JWT dan mengembalikan klaimnya.
func ParseToken(secret, tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
