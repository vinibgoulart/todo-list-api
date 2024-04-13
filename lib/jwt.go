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

func JWTVerifyToken(tokenString string) (*jwt.Token, error) {
	envFile, _ := godotenv.Read(".env")

	secretKey := []byte(envFile["JWT_KEY"])

	tokenString = strings.Split(tokenString, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
