package tests

import (
	"goPractice/task_manager/infrastructure"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func init() {
	log.Println("Token Expiration", godotenv.Load("../../.env"), os.Getenv("JWT_EXP"))
}

// test the passowrd hashing function
func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := infrastructure.HashPassword(password)
	require.NoError(t, err, "Failed to hash password")
	require.NotEqual(t, password, hashedPassword, "Password and hashed password should be different")

}

func TestCheckPasswordHash(t *testing.T) {

	password := "password"
	hashedPassword, _ := infrastructure.HashPassword(password)
	require.True(t, infrastructure.CheckPasswordHash(password, hashedPassword), "Password and hashed password should match")
}

func TestGenerateToken(t *testing.T) {

	token, err := infrastructure.GenerateToken("1", "admin")
	log.Println(token)
	require.NoError(t, err, "Failed to generate token")
	require.NotEmpty(t, token, "Token should not be empty")
}

// func TestParseToken(t *testing.T) {
// 	token, err := infrastructure.GenerateToken("1", "admin")
// 	require.NoError(t, err, "Failed to generate token")

// 	claims, err := infrastructure.ParseToken("Bearer " + token)
// 	require.NoError(t, err, "Failed to parse token")
// 	require.NotEmpty(t, claims, "Claims should not be empty")
// }
// lets create PR
