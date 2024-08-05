package main

import (
	"goPractice/task_manager/config"
	"goPractice/task_manager/controllers"
	"goPractice/task_manager/data"
	"goPractice/task_manager/router"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Trainer struct {
	Name string `bson:"name,omitempty" json:"name"`
	Age  int    `bson:"age,omitempty" json:"age"`
	City string `bson:"city,omitempty" json:"city"`
}
type Server struct {
	Router *gin.Engine
	DB     *mongo.Database
	Port   string
}

func (s *Server) Run() {
	log.Println("Starting server on port " + s.Port)
	s.Router = gin.Default()

	//force color tobe displayed in console
	gin.ForceConsoleColor()
	s.DB = config.NewMongoClient().Database("test")
	api := s.Router.Group("/api")

	task_service := data.NewTaskService(s.DB.Collection("tasks"))
	task_controller := controllers.NewTaskController(task_service)
	router.SetUpTaskRouter(api, task_controller)

	user_service := data.NewUserService(s.DB.Collection("users"))
	user_controller := controllers.NewUserController(user_service)
	router.SetUpUserRouter(api, user_controller)
	if s.Port == "" {
		s.Port = "8080"
	}
	s.Router.Run(":" + s.Port)

}
func main() {

	server := Server{Port: "8080"}
	server.Run()
}
