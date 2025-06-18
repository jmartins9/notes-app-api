package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmartins9/notes-app-api/controllers"
)

func AuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
}
