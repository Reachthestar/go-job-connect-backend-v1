package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey = os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY is not set in .env file")
	}
}

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"UserId": userId,
		"exp": time.Now().Add(time.Hour * 10).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64,error){
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
		return nil, errors.New("Unexpected signing method")
	}

	return []byte(secretKey), nil
	})	

		if err != nil {
		return 0, errors.New("Could not parse token.")
	}

	isTokenValid := parsedToken.Valid

	if !isTokenValid {
		return 0, errors.New("Invalid token!")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		errors.New("Invalid token claims")
	}

	userId := int64(claims["userId"].(float64))
	fmt.Println("userId from JWT", userId)

	return userId, nil
}