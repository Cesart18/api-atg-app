package controllers

import (
	"net/http"
	"strings"

	"github.com/cesart18/ranking_app/db"
	"github.com/cesart18/ranking_app/models"
	"github.com/cesart18/ranking_app/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) Signup(c *gin.Context) {

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if c.BindJSON(&body) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}

	msg, err := uc.UserService.Signup(body.Username, body.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, msg)
}

func (uc *UserController) Login(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if c.BindJSON(&body) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	token, err := uc.UserService.Login(body.Username, body.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) Logout(c *gin.Context) {

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

	t := models.RevokedToken{
		Token: tokenString,
	}
	msg, err := uc.UserService.Logout(t)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": msg})

}

func (uc *UserController) Validate(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "usuario autorizado",
	})
}
