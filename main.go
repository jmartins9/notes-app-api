package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmartins9/notes-app-api/routes"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		routes.UsersRoutes(api)
		routes.AuthRoutes(api)
		routes.TasksRoutes(api)
		routes.SessionsRoutes(api)
	}

	r.Run(":8080")
}
