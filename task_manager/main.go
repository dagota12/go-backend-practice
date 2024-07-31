package main

import (
	"fmt"
	"net/http"
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
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"tasks": tasks})
	})

	fmt.Println("Task Manager-API GIN")

	router.Run(":8080")
}
