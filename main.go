package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmartins9/notes-app-api/controllers"
	"github.com/jmartins9/notes-app-api/models"
	"github.com/jmartins9/notes-app-api/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=notesdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate your models
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Inject DB into controllers
	controllers.SetDatabase(db)

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
