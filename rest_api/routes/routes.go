package routes

import (
	"github.com/gin-gonic/gin"
	"repos/project/controllers"
	"repos/project/dao"
	"repos/project/service"
)

func SetupRouter() *gin.Engine {

	userDao := dao.NewUserDaoImpl()
	userService := service.NewUserServiceImpl(userDao)
	userContoller := controllers.NewUserController(userService)

	r := gin.Default()

	r.GET("user", userContoller.GetUser)
	r.POST("user", userContoller.CreateUser)
	r.PUT("user/:id", userContoller.UpdateUser)
	r.GET("user/:id", userContoller.GetUserById)

	return r
}
