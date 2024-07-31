package router

import (
	"goPractice/task_manager/controllers"
	"goPractice/task_manager/data"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {
	api := r.Group("/api")

	task_service := data.NewTaskService()
	task_controller := controllers.NewTaskController(task_service)
	taskGroup := api.Group("/task")
	taskGroup.GET("/", task_controller.GetTasks)
	taskGroup.POST("/", task_controller.CreateTask)
	taskGroup.GET("/:id", task_controller.GetTask)
	taskGroup.PUT("/:id", task_controller.UpdateTask)
	taskGroup.DELETE("/:id", task_controller.DeleteTask)
	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong"})
	})

}
