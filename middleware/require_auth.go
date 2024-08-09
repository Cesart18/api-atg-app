package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cesart18/ranking_app/config"
	"github.com/cesart18/ranking_app/db"
	"github.com/cesart18/ranking_app/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {

	authHeader := c.GetHeader("authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Usuario no autenticado")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	var revokedToken models.RevokedToken
	db.DB.Find(&revokedToken, "token = ?", tokenString)

	if revokedToken.ID != 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Usuario no autenticado")
	}
	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing methos: %v", t.Header)
		}
		return []byte(config.Secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Usuario no autenticado")
		}
		id := claims["sub"]
		var user models.User
		db.DB.First(&user, id)

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Usuario no autenticado")
		}

		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Usuario no autenticado")
	}
}
