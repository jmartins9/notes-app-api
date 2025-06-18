package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTasks(c *gin.Context) {
	tasks := []gin.H{
		{"id": 1, "title": "Estudar Go", "done": false},
		{"id": 2, "title": "Fazer API com Gin", "done": true},
	}

	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task struct {
		Title string `json:"title"`
		Done  bool   `json:"done"`
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tarefa criada com sucesso",
		"task":    task,
	})
}
