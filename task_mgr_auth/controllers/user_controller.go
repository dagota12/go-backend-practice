package controllers

import (
	"goPractice/task_manager/data"
	"goPractice/task_manager/middleware"
	"goPractice/task_manager/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService data.UserService
}

func NewUserController(service data.UserService) *UserController {
	return &UserController{userService: service}
}
func (uc *UserController) GetUsers(c *gin.Context) {

	users, err := uc.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting all users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
func (uc *UserController) CreateUser(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login faild Error while Parsing the request body"})
		return
	}
	//check if user aready exist
	existingUser, _ := uc.userService.FilterUser(user.Username)

	if existingUser.Username != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		return
	}
	//hash password using bcrypt
	hashedPwd, err := middleware.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing the password"})
		return
	}
	user.Password = hashedPwd
	//create the user in the database
	//check if database is empty or if there are no users this user role becomes admin
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting all users"})
		return
	}
	if len(users) == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	newUser, err := uc.userService.CreateUser(user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while Creating User"})
		return
	}
	// c.JSON(http.StatusCreated, gin.H{"message": "create user"})
	c.JSON(http.StatusCreated, gin.H{"id": newUser.ID, "username": newUser.Username}) //new user created

}

func (uc *UserController) UserLogin(c *gin.Context) {

	var data models.LoginForm

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.userService.FilterUser(data.Username)
	log.Println("user Logging in:", user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No such User found"})
		return
	}
	//check password
	if !middleware.CheckPasswordHash(data.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
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
func (uc *UserController) PromoteUser(c *gin.Context) {

	id := c.Param("id")

	if _, err := uc.userService.PromoteUser(id); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while promoting user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "admin promoted"})
}
