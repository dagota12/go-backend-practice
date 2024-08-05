package router

import (
	"goPractice/task_manager/controllers"
	"goPractice/task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpTaskRouter(r *gin.RouterGroup, tc *controllers.TaskController) {
	{
		taskGroup := r.Group("/task")
		taskGroup.GET("/", middleware.Authorize("user", "admin"), tc.GetTasks)
		taskGroup.POST("/", middleware.Authorize("admin"), tc.CreateTask)
		taskGroup.GET("/:id", middleware.Authorize("user", "admin"), tc.GetTask)
		taskGroup.PUT("/:id", middleware.Authorize("admin"), tc.UpdateTask)
		taskGroup.DELETE("/:id", middleware.Authorize("admin"), tc.DeleteTask)
	}
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong"})
	})

}
func SetUpUserRouter(r *gin.RouterGroup, uc *controllers.UserController) {

	userGroup := r.Group("/user")
	// userGroup.GET("/", uc.GetUsers)
	userGroup.POST("/register", uc.CreateUser)
	userGroup.POST("/login", uc.UserLogin)

	adminGroup := r.Group("/admin")

	adminGroup.POST("register", uc.CreateAdmin)
	adminGroup.POST("/login", uc.AdminLogin)

}
