package controllers

import (
	"goPractice/task_manager/data"
	"goPractice/task_manager/middleware"
	"goPractice/task_manager/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login Error while Parsing the request"})
		return
	}
	//check for if existing user if yes send conflict
	existingUser, _ := uc.userService.FilterUser(bson.M{"username": user.Username})

	if existingUser.Username != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	newUser, err := uc.userService.CreateUser(user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while Creating User"})
		return
	}
	// c.JSON(http.StatusCreated, gin.H{"message": "create user"})
	c.JSON(http.StatusCreated, newUser) //new user created

}

func (uc *UserController) UserLogin(c *gin.Context) {

	var data models.LoginForm

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.userService.FilterUser(bson.M{"username": data.Username, "password": data.Password})
	log.Println("user Logging in:", user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No such User found"})
		return
	}
	//generate jwt token
	token, _ := middleware.GenerateToken(user.ID.Hex(), user.Role)
	log.Println("token:", token)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) CreateAdmin(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "create admin"})
}
func (uc *UserController) AdminLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "admin login"})
}
