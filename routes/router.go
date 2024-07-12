package routes

import (
	"github.com/cesart18/ranking_app/controllers"
	"github.com/cesart18/ranking_app/middleware"
	"github.com/cesart18/ranking_app/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	playerService := services.PlayerService{}
	playerController := controllers.NewPlayerController(&playerService)

	userService := services.UserService{}
	userController := controllers.NewUserController(&userService)

	api := r.Group("/api")
	{
		api.POST("/player", middleware.RequireAuth, playerController.CreatePlayer)
		api.GET("/players", playerController.GetPlayers)
		api.GET("/player/:id", playerController.GetPlayerById)
		api.PATCH("/player/:id", middleware.RequireAuth, playerController.UpdatePlayer)
		api.POST("/double_point/:id", middleware.RequireAuth, playerController.AddDoublePoint)
		api.POST("/single_point/:id", middleware.RequireAuth, playerController.AddSinglePoint)
		api.DELETE("/player/:id", middleware.RequireAuth, playerController.DeletePlayer)
	}
	auth := r.Group("/auth")
	{
		auth.POST("/signup", middleware.RequireAdminSignup, userController.Signup)
		auth.POST("/login", userController.Login)
		auth.POST("/logout", userController.Logout)
		auth.GET("/validate", middleware.RequireAuth, userController.Validate)
	}

	return r
}
