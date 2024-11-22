package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(userId string, role string) (string, error) {
	key := os.Getenv("SECRETKEY")
	var secretKey = []byte(key)
	fmt.Println(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"role":   role,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	fmt.Println(err)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
