package lib

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func JWTCreateToken(identifier string) (string, error) {
	envFile, _ := godotenv.Read(".env")

	secretKey := []byte(envFile["JWT_KEY"])

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"identifier": identifier,
			"exp":        time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func UserVerifyToken(tokenString string) error {
	envFile, _ := godotenv.Read(".env")

	secretKey := envFile["JWT_KEY"]

	tokenString = strings.Split(tokenString, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
