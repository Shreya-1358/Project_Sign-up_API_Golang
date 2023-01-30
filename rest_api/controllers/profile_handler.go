package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"repos/project/model"
	"repos/project/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(us service.UserService) *UserController {
	return &UserController{userService: us}
}

func (uc *UserController) GetUser(c *gin.Context) {

	user, err := uc.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Records not Found"})

	} else {
		c.JSON(http.StatusOK, user)
	}
}
func (uc *UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.userService.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Employee Not Exist"})

	} else {
		c.JSON(http.StatusOK, user)
	}
}
func (uc *UserController) CreateUser(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating Employee"})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (uc *UserController) UpdateUser(c *gin.Context) {

	var user model.User
	id := c.Param("id")
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.userService.UpdateUser(&user, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Employee Not Exist"})
	} else {
		c.JSON(http.StatusOK, user)
	}

}
