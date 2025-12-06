package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SecretKey = "secret"

func GenerateToken(firstName, lastName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"firstName": firstName,
		"lastName":  lastName,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(SecretKey))
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
