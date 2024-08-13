package mocks

import (
	"goPractice/task_manager/infrastructure"

	"github.com/gin-gonic/gin"
)

func Authorize(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims := infrastructure.Claims{
			UserID: "1",
			Role:   role,
		}
		ctx.Set("claim", &claims)

	}
}
