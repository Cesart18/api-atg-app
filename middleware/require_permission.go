package middleware

import (
	"net/http"

	"github.com/cesart18/ranking_app/config"
	"github.com/gin-gonic/gin"
)

func RequireAdminSignup(c *gin.Context) {
	// Obtener la clave personal del administrador
	adminKey := c.GetHeader("key")

	// Verificar si la clave personal es válida
	if adminKey != config.AdminKey {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "No tienes permisos para crear usuarios",
		})
		return
	}

	// Si la clave es válida, permitir que la solicitud continúe
	c.Next()
}
