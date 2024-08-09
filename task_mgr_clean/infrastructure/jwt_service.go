package infrastructure

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userID string, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(EXP_TIME)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(JWT_SECRET))
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
		return []byte(JWT_SECRET), nil
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
