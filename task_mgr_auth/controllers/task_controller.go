package controllers

import (
	"errors"
	"goPractice/task_manager/data"
	"goPractice/task_manager/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService data.TaskService
}

func NewTaskController(taskService data.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

// get current user claims from the middleware
func getClaims(c *gin.Context) (map[string]string, error) {
	user_id, ok1 := c.Get("user_id")
	role, ok2 := c.Get("role")
	log.Printf("CURRENT USER: user_id: %v, role: %v", user_id, role)

	//both fields exist
	if !ok1 || !ok2 {
		return nil, errors.New("unauthorized")
	}
	data := map[string]string{"user_id": user_id.(string), "role": role.(string)}
	return data, nil
}

// create task
func (tc *TaskController) CreateTask(c *gin.Context) {
	// data, err := getClaims(c)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	// 	return
	// }
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	//extract the user id from context map
	// if task.UserId, err = primitive.ObjectIDFromHex(data["user_id"]); err != nil {
	// 	log.Println(err.Error())
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error,be sure to log in"})
	// 	return
	// }
	newTask, err := tc.taskService.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

func (tc *TaskController) GetTasks(c *gin.Context) {

	tasks := tc.taskService.GetTasks()

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetUserTasks(c *gin.Context) {
	//get the user claims thats sent from middleware
	data, err := getClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	tasks := tc.taskService.GetUserTasks(data["user_id"])

	c.JSON(http.StatusOK, tasks)
}
func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.taskService.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad payload," + err.Error()})
		return
	}

	updatedTask, err := tc.taskService.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := tc.taskService.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
