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

// PlayerController godoc
// @Summary      Create a new player
// @Description  Create a new player
// @Tags         players
// @Accept       json
// @Produce      json
// @Param        player body models.Player true "Player data"
// @Success      200  {object}  models.Player
// @Failure      400  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /players [post]

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

// PlayerController godoc
// @Summary      Get all players
// @Description  Get a list of all players
// @Tags         players
// @Produce      json
// @Success      200  {array}   models.Player
// @Failure      400  {object}  gin.H
// @Router       /players [get]

func (uc *PlayerController) GetPlayers(c *gin.Context) {
	users, err := uc.PlayerController.GetPlayers()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

// PlayerController godoc
// @Summary      Get a player by ID
// @Description  Get a player by their ID
// @Tags         players
// @Produce      json
// @Param        id   path      int  true  "Player ID"
// @Success      200  {object}  models.Player
// @Failure      500  {object}  gin.H
// @Router       /players/{id} [get]
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

// PlayerController godoc
// @Summary      Update a player
// @Description  Update a player's name
// @Tags         players
// @Produce      json
// @Param        id    path      int    true  "Player ID"
// @Param        name  query     string true  "Player name"
// @Success      200  {object}  gin.H
// @Failure      400  {object}  gin.H
// @Router       /players/{id} [put]
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

// PlayerController godoc
// @Summary      Add double points to a player
// @Description  Add double points to a player's score
// @Tags         players
// @Produce      json
// @Param        id     path      int    true  "Player ID"
// @Param        points query     int    true  "Points to add"
// @Success      200  {object}  models.Player
// @Failure      400  {object}  gin.H
// @Router       /players/{id}/double-points [post]
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

// PlayerController godoc
// @Summary      Add single points to a player
// @Description  Add single points to a player's score
// @Tags         players
// @Produce      json
// @Param        id     path      int    true  "Player ID"
// @Param        points query     int    true  "Points to add"
// @Success      200  {object}  models.Player
// @Failure      400  {object}  gin.H
// @Router       /players/{id}/single-points [post]
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

// PlayerController godoc
// @Summary      Delete a player
// @Description  Delete a player by their ID
// @Tags         players
// @Produce      json
// @Param        id   path      int  true  "Player ID"
// @Success      200  {object}  models.Player
// @Failure      400  {object}  gin.H
// @Router       /players/{id} [delete]

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
