package controllers

import (
	"goPractice/task_manager/data"
	"goPractice/task_manager/models"
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

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser, err := uc.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "create user"})
	c.JSON(http.StatusCreated, newUser) //new user created

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
