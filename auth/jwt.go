package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("my_secret_key")

func GenerateJWT(userID uint, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"userID":   userID,
		"username": username,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
