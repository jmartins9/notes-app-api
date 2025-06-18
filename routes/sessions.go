package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmartins9/notes-app-api/controllers"
)

func SessionsRoutes(rg *gin.RouterGroup) {
	sessions := rg.Group("/sessions")
	{
		sessions.GET("/", controllers.GetSessions)
		sessions.DELETE("/:id", controllers.DeleteSession)
	}
}
