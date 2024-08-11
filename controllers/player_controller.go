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

type PlayerIDsRequest struct {
	PlayerIDs []int  `json:"ids"`
	Score     string `json:"score"`
}
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
		log.Println(err)
	}

	player, err := uc.PlayerController.GetPlayerById(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, player)
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

func (uc *PlayerController) ToggleMembership(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}
	msg, err := uc.PlayerController.ToggleMembership(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	c.IndentedJSON(http.StatusOK, msg)
}

func (uc *PlayerController) TogglePayedBalls(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}
	msg, err := uc.PlayerController.TogglePayedBalls(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	c.IndentedJSON(http.StatusOK, msg)
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

func (uc *PlayerController) AddMatch(c *gin.Context) {

	var request PlayerIDsRequest

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	msg, err := uc.PlayerController.AddMatch(request.PlayerIDs, request.Score)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, msg)
}
