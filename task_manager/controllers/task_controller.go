package controllers

import (
	"goPractice/task_manager/data"
	"goPractice/task_manager/models"
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
func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

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
