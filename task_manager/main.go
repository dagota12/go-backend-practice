package main

import (
	"goPractice/task_manager/controllers"
	"goPractice/task_manager/data"
	"goPractice/task_manager/router"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

var tasks = map[string]Task{
	"1": {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	"2": {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	"3": {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func main() {
	router_inst := gin.Default()
	task_service := data.NewTaskService()
	// task_service.Tasks = tasks

	task_controller := controllers.NewTaskController(task_service)
	router.SetUpRouter(router_inst, task_controller)

	router_inst.Run(":8080")
}
