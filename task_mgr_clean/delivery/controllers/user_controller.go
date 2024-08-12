package controllers

import (
	"goPractice/task_manager/domain"
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

	newUser, err := uc.userUsecase.CreateUser(user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	token, err := uc.userUsecase.Login(data)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) PromoteUser(c *gin.Context) {

	id := c.Param("id")

	if err := uc.userUsecase.PromoteUser(id); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error while promoting user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user promoted success!"})
}
