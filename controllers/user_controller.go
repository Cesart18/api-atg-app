package controllers

import (
	"net/http"

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

// UserController godoc
// @Summary      Sign up a new user
// @Description  Create a new user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.User true "User data"
// @Success      200  {object}  gin.H
// @Failure      400  {object}  gin.H
// @Router       /signup [post]
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

// UserController godoc
// @Summary      Login a user
// @Description  Authenticate a user and get a JWT token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.User true "User credentials"
// @Success      200  {object}  gin.H
// @Failure      400  {object}  gin.H
// @Router       /login [post]
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
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24, "", "", false, true)

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}

// UserController godoc
// @Summary      Logout a user
// @Description  Revoke the user's JWT token and log them out
// @Tags         users
// @Produce      json
// @Success      200  {object}  gin.H
// @Failure      409  {object}  gin.H
// @Router       /logout [post]
func (uc *UserController) Logout(c *gin.Context) {

	token, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusConflict)
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

	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"message": msg})

}

// UserController godoc
// @Summary      Validate a user
// @Description  Validate the user's authentication status
// @Tags         users
// @Produce      json
// @Success      200  {object}  gin.H
// @Router       /validate [get]
func (uc *UserController) Validate(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "usuario autorizado",
	})
}
