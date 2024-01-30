package jwt

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id, role string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  jwt.NewNumericDate(time.Now().Add(15 * time.Minute)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

func ExtractToken(bearerToken string) string {

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func VerifyToken(token string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	tokens, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := tokens.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to get claims from token")
	}

	if !tokens.Valid {
		return "", errors.New("token is not valid")
	}

	expires, ok := claims["exp"].(float64)
	if !ok {
		return "", errors.New("failed to get expiration time from token")
	}

	expiresTime := time.Unix(int64(expires), 0)
	if time.Now().After(expiresTime) {
		return "", errors.New("token has expired")
	}

	userId, ok := claims["id"].(string)
	if !ok || userId == "" {
		return "", errors.New("failed to get user ID from token")
	}

	return userId, nil
}
