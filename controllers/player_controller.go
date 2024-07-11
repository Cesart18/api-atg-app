package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/cesart18/ranking_app/models"
	"github.com/cesart18/ranking_app/services"
	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	UserService services.PlayerServiceInterface
}

func NewUserController(userService services.PlayerServiceInterface) *PlayerController {
	return &PlayerController{
		UserService: userService,
	}
}

func (uc *PlayerController) CreatePlayer(c *gin.Context) {

	var user models.Player

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	user.Name = strings.TrimSpace(user.Name)
	msg, err := uc.UserService.CreatePlayer(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, msg)

}

func (uc *PlayerController) GetPlayers(c *gin.Context) {
	users, err := uc.UserService.GetPlayers()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (uc *PlayerController) GetPlayerById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	user, err := uc.UserService.GetPlayerById(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}
func (uc *PlayerController) UpdatePlayer(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}
	name := c.Query("name")

	msg, err := uc.UserService.UpdatePlayer(id, name)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, msg)
}
func (uc *PlayerController) AddDoublePoint(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	pointParam := c.Query("points")
	p, err := strconv.Atoi(pointParam)
	if err != nil {
		log.Fatal(err)
	}
	user, err := uc.UserService.AddDoublePoint(id, p)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}
func (uc *PlayerController) AddSinglePoint(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}
	pointParam := c.Query("points")
	p, err := strconv.Atoi(pointParam)
	if err != nil {
		log.Fatal(err)
	}

	user, err := uc.UserService.AddSinglePoint(id, p)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}
func (uc *PlayerController) DeletePlayer(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	user, err := uc.UserService.DeletePlayer(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}
