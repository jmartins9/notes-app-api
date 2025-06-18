package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmartins9/notes-app-api/controllers"
)

func UsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.GET("/", controllers.GetUsers)
		users.POST("/", controllers.CreateUser)
	}
}
