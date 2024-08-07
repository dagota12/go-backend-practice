package router

import (
	"goPractice/task_manager/controllers"
	"goPractice/task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpTaskRouter(r *gin.RouterGroup, tc *controllers.TaskController) {
	{
		taskGroup := r.Group("/task")
		taskGroup.GET("/", middleware.Authorize("user", "admin"), tc.GetUserTasks)
		taskGroup.GET("/all", middleware.Authorize("admin"), tc.GetTasks)

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
	//public
	userGroup.POST("/register", uc.CreateUser)
	userGroup.POST("/login", uc.UserLogin)
	//protected
	protected := userGroup.Group("/")
	protected.Use(middleware.Authorize("admin"))

	protected.GET("/", uc.GetUsers) //get all users
	protected.POST("/promote/:id", uc.PromoteUser)

	// adminGroup := r.Group("/admin")

	// adminGroup.POST()

	// adminGroup.POST("register", uc.CreateAdmin)
	// adminGroup.POST("/login", uc.AdminLogin)

}
