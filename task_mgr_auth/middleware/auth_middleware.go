package middleware

import (
	"fmt"
	"goPractice/task_manager/config"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// struct for the payload in the JWT token
type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func Authorize(allowedRoles ...string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
			ctx.Abort()
			return
		}
		//Parse the JWT token
		claim, err := ParseToken(authHeader)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		for _, role := range allowedRoles {
			if claim["role"] == role {
				log.Printf("CLAIM: user_id: %v, role: %v", claim["user_id"], claim["role"])
				ctx.Set("user_id", claim["user_id"])
				ctx.Set("role", claim["role"])
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Insufficient permissions"})
		ctx.Abort()

	}

}

func ParseToken(authHeader string) (jwt.MapClaims, error) {
	authParts := strings.Split(authHeader, " ")

	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(config.JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims) //type assertion for jwt.MapClaims
	fmt.Println("Claims: ", claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid JWT token")
	}

	return claims, nil
}

func GenerateToken(userID string, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_SECRET))
}
func HashPassword(password string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPwd), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println("comparing hash and password:", err.Error())
	}
	return err == nil
}
