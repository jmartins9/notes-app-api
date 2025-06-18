package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmartins9/notes-app-api/controllers"
)

func TasksRoutes(rg *gin.RouterGroup) {
	tasks := rg.Group("/tasks")
	{
		tasks.GET("/", controllers.GetTasks)
		tasks.POST("/", controllers.CreateTask)
	}
}
