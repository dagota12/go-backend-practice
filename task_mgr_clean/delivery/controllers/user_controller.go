package controllers

import (
	"goPractice/task_manager/domain"
	"goPractice/task_manager/infrastructure"
	"goPractice/task_manager/usecases"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecases.UserUsecase
}

func NewUserController(usecase usecases.UserUsecase) *UserController {
	return &UserController{
		userUsecase: usecase,
	}
}

func (uc *UserController) GetUsers(c *gin.Context) {

	users := uc.userUsecase.GetUsers()

	c.JSON(http.StatusOK, users)
}

// create user
func (uc *UserController) CreateUser(c *gin.Context) {

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login faild Error while Parsing the request body"})
		return
	}
	//check if user aready exist
	existingUser, _ := uc.userUsecase.FilterUser(user.Username)

	if existingUser.Username != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		return
	}
	//hash password using bcrypt
	hashedPwd, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing the password"})
		return
	}
	user.Password = hashedPwd
	//create the user in the database
	//check if database is empty or if there are no users this user role becomes admin
	users := uc.userUsecase.GetUsers()
	if len(users) == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	newUser, err := uc.userUsecase.CreateUser(user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while Creating User"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": newUser.ID, "username": newUser.Username}) //new user created

}

//user login

func (uc *UserController) UserLogin(c *gin.Context) {

	var data domain.LoginForm

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.userUsecase.FilterUser(data.Username)
	log.Println("user Logging in:", user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No such User found"})
		return
	}
	//check password
	valid := infrastructure.CheckPasswordHash(data.Password, user.Password)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password doesn't match"})
		return
	}
	//generate jwt token
	token, _ := infrastructure.GenerateToken(user.ID.Hex(), user.Role)
	// log.Println("token:", token)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) PromoteUser(c *gin.Context) {

	id := c.Param("id")

	if err := uc.userUsecase.PromoteUser(id); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while promoting user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user promoted success!"})
}
