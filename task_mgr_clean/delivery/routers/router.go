package routers

import (
	"goPractice/task_manager/delivery/controllers"
	"goPractice/task_manager/infrastructure"
	"goPractice/task_manager/repositories"
	"goPractice/task_manager/usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup() {
	client := infrastructure.NewMongoClient()
	db := client.Database(infrastructure.DB_NAME)
	router := gin.Default()
	api := router.Group("/api")
	UserRouter(db, api)
	TaskRouter(db, api)

	router.Run(":8080")
}
func UserRouter(db *mongo.Database, r *gin.RouterGroup) {
	userRepo := repositories.NewUserRepository(db)
	uc := controllers.NewUserController(usecases.NewUserUsecase(userRepo))
	group := r.Group("/user")
	//public
	group.POST("/register", uc.CreateUser)
	group.POST("/login", uc.UserLogin)
	//protected
	protected := group.Group("/")
	protected.Use(infrastructure.Authorize("admin"))

	protected.GET("/", uc.GetUsers) //get all users
	protected.POST("/promote/:id", uc.PromoteUser)
}
func TaskRouter(db *mongo.Database, router *gin.RouterGroup) {
	tasksRepo := repositories.NewTaskRepository(db)
	tc := controllers.NewTaskController(usecases.NewTaskUsecase(tasksRepo))
	taskG := router.Group("/task")
	taskG.Use(infrastructure.Authorize("user", "admin"))

	taskG.GET("/", tc.GetUserTasks)
	taskG.GET("/:id", tc.GetTask)

	//protected
	adminOnly := taskG.Group("/")
	adminOnly.Use(infrastructure.Authorize("admin"))

	adminOnly.GET("/all", tc.GetTasks)
	adminOnly.POST("/", tc.CreateTask)
	adminOnly.PUT("/:id", tc.UpdateTask)
	adminOnly.DELETE("/:id", tc.DeleteTask)

}
