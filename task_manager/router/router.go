package router

import (
	"goPractice/task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine, tc *controllers.TaskController) {
	api := r.Group("/api")
	{
		taskGroup := api.Group("/task")
		taskGroup.GET("/", tc.GetTasks)
		taskGroup.POST("/", tc.CreateTask)
		taskGroup.GET("/:id", tc.GetTask)
		taskGroup.PUT("/:id", tc.UpdateTask)
		taskGroup.DELETE("/:id", tc.DeleteTask)
	}
	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong"})
	})

}
