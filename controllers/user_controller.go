package controllers

import (
	"net/http"
	"os"

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
	domain := os.Getenv("COOKIE_DOMAIN")
	if domain == "" {
		domain = "" // Dejar vacío para permitir acceso desde cualquier dominio
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24, "/", domain, false, true)

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) Logout(c *gin.Context) {

	token, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		return
	}
	t := models.RevokedToken{
		Token: token,
	}
	msg, err := uc.UserService.Logout(t)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		return
	}
	domain := os.Getenv("COOKIE_DOMAIN")
	if domain == "" {
		domain = "" // Dejar vacío para permitir acceso desde cualquier dominio
	}

	c.SetCookie("Authorization", "", -1, "/", domain, false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"message": msg})

}

func (uc *UserController) Validate(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "usuario autorizado",
	})
}
