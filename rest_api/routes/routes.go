package routes

import (
	"github.com/gin-gonic/gin"
	"repos/project/controllers"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	//grp1 := r.Group("/user-api")
	r.GET("/user", controllers.GetUsers)
	r.POST("/user", controllers.CreateUser)
	r.GET("user/:id", controllers.GetUserByID)
	//grp1.GET("/user", controllers.GetUsers)
	//grp1.PUT("user/:id", controllers.UpdateUser)
	//grp1.DELETE("user/:id", controllers.DeleteUser)

	return r
}
