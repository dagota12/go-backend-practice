package controllers

import (
	"errors"
	"goPractice/task_manager/domain"
	"goPractice/task_manager/infrastructure"
	"goPractice/task_manager/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUsecase usecases.TaksUsecase
}

func NewTaskController(usecase usecases.TaksUsecase) *TaskController {
	return &TaskController{taskUsecase: usecase}
}

// get current user claims from the middleware
func getClaims(c *gin.Context) (map[string]string, error) {
	claims, ok := c.Get("claim")

	//if claims does not exist, return unauthorized
	if !ok {
		return nil, errors.New("unauthorized")
	}
	user := claims.(*infrastructure.Claims)
	data := map[string]string{"user_id": user.UserID, "role": user.Role}
	return data, nil
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	newTask, err := tc.taskUsecase.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

func (tc *TaskController) GetTasks(c *gin.Context) {

	tasks := tc.taskUsecase.GetTasks()

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetUserTasks(c *gin.Context) {
	//get the user claims thats sent from middleware
	data, err := getClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	tasks := tc.taskUsecase.GetUserTasks(data["user_id"])

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.taskUsecase.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad payload," + err.Error()})
		return
	}

	updatedTask, err := tc.taskUsecase.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := tc.taskUsecase.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
