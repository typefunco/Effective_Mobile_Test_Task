package utils

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var JwtSecret []byte

func initSecret() error {
	err := godotenv.Load("./config/.env")
	if err != nil {
		slog.Info("Error loading .env file [ERROR]")
		return err
	}

	jwtEnv, ok := os.LookupEnv("DBURL")
	if !ok {
		slog.Info("jwt not found in environment [ERROR]")
		return err
	}
	var jwtSec = []byte(jwtEnv)
	JwtSecret = jwtSec
	return nil
}

func GetSecret() []byte {
	secret := JwtSecret
	return secret
}

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1000).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func GetUsernameFromJWT(tokenString string) (string, error) {
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return "", err
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username claim is missing or invalid")
	}

	return username, nil
}
