package helper

import (
	"time"
	"zaxx/backend/config"
	
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECERT", ""))

func GenerateToken(username  string)string {
	expirateTIme := time.Now().Add(60 * time.Minute)

	claims := &jwt.RegisteredClaims{
		Subject: username,
		ExpiresAt: jwt.NewNumericDate(expirateTIme),
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256,claims).SignedString(jwtKey)

	return token	
}