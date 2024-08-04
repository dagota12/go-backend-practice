package controllers

import (
	"goPractice/task_manager/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService data.UserService
}

func NewUserController(service data.UserService) *UserController {
	return &UserController{userService: service}
}
func (uc *UserController) CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "create user"})

}

func (uc *UserController) UserLogin(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "user login"})
}

func (uc *UserController) CreateAdmin(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "create admin"})
}
func (uc *UserController) AdminLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "admin login"})
}
