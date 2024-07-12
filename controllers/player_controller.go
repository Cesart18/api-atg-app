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
	PlayerController services.PlayerServiceInterface
}

func NewPlayerController(playerController services.PlayerServiceInterface) *PlayerController {
	return &PlayerController{
		PlayerController: playerController,
	}
}

func (uc *PlayerController) CreatePlayer(c *gin.Context) {

	var user models.Player

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	user.Name = strings.TrimSpace(user.Name)
	msg, err := uc.PlayerController.CreatePlayer(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, msg)

}

func (uc *PlayerController) GetPlayers(c *gin.Context) {
	users, err := uc.PlayerController.GetPlayers()

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

	user, err := uc.PlayerController.GetPlayerById(id)

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

	msg, err := uc.PlayerController.UpdatePlayer(id, name)

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
	user, err := uc.PlayerController.AddDoublePoint(id, p)

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

	user, err := uc.PlayerController.AddSinglePoint(id, p)

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

	user, err := uc.PlayerController.DeletePlayer(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}
