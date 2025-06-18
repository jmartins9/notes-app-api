package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSessions(c *gin.Context) {
	sessions := []gin.H{
		{"id": "s1", "user": "admin@example.com", "active": true},
		{"id": "s2", "user": "user@example.com", "active": false},
	}

	c.JSON(http.StatusOK, sessions)
}

func DeleteSession(c *gin.Context) {
	sessionID := c.Param("id")

	// Aqui normalmente você removeria a sessão do banco de dados ou cache
	c.JSON(http.StatusOK, gin.H{"message": "Sessão removida com sucesso", "sessionId": sessionID})
}
