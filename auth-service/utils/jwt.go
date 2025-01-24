package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_ACCESS_SECRET")
	if secret == "" {
		return "", errors.New("JWT_ACCESS_SECRET_SECRET is not set")
	}
	log.Println(secret)
	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken() string {
	return uuid.New().String()
}

// func VerifyToken(tokenString string) (uint, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, jwt.ErrSignatureInvalid
// 		}
// 		return os.Getenv("JWT_SECRET"), nil
// 	})
// 	if err != nil {
// 		return 0, err
// 	}
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		userID := uint(claims["user_id"].(float64))
// 		return userID, nil
// 	}
// 	return 0, jwt.ErrSignatureInvalid
// }
