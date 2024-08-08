package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
			if claim.Role == role {
				// log.Printf("CLAIM: user_id: %v, role: %v", claim["user_id"], claim["role"])
				ctx.Set("claim", claim)
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Insufficient permissions"})
		ctx.Abort()

	}

}
