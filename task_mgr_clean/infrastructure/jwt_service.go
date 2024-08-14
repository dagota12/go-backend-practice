package infrastructure

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, role string) (string, error) {
	exp_time, _ := strconv.Atoi(os.Getenv("JWT_EXP"))
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(exp_time))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseToken(authHeader string) (Claims, error) {
	authParts := strings.Split(authHeader, " ")

	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return Claims{}, fmt.Errorf("invalid authorization header format")
	}
	claims := Claims{}
	token, err := jwt.ParseWithClaims(authParts[1], &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return Claims{}, err
	}

	fmt.Println("Claims: ", claims)
	if !token.Valid {
		return Claims{}, fmt.Errorf("invalid JWT token")
	}

	return claims, nil
}
