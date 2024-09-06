package services

import (
	"Project-Management-System/models"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtSecret = []byte("secret-key")

func GenerateJWT(user models.UserModel) (string, error) {
	claims := jwt.MapClaims{
		"userId":   user.UserId,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}
