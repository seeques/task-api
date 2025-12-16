package auth

import (
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	ID int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int, secret string) (string, error) {
	claims := CustomClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(secret))

	return ss, err
}

func ValidateToken(tokenString, secret string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	return claims.ID, nil
}