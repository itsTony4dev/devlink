package utils

import (
	"time" 
	"devlink/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.GetEnv("JWT_SECRET", "secret")) 

// GetJWTSecret returns the JWT secret key
func GetJWTSecret() []byte {
	return jwtSecret
}

func GenerateJWT(userID uint, email string, usename string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"email":    email,
		"username": usename,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 3 days expiry
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
