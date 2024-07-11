package routes

import (
	"github.com/cesart18/ranking_app/controllers"
	"github.com/cesart18/ranking_app/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	service := services.PlayerService{}
	controller := controllers.NewUserController(&service)

	api := r.Group("/api")
	{
		api.POST("/player", controller.CreatePlayer)
		api.GET("/players", controller.GetPlayers)
		api.GET("/player/:id", controller.GetPlayerById)
		api.PATCH("/player/:id", controller.UpdatePlayer)
		api.POST("/double_point/:id", controller.AddDoublePoint)
		api.POST("/single_point/:id", controller.AddSinglePoint)
		api.DELETE("/player/:id", controller.DeletePlayer)
	}

	return r
}
