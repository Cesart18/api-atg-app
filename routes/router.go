package routes

import (
	"time"

	"github.com/cesart18/ranking_app/controllers"
	"github.com/cesart18/ranking_app/middleware"
	"github.com/cesart18/ranking_app/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))

	playerService := services.PlayerService{}
	playerController := controllers.NewPlayerController(&playerService)

	userService := services.UserService{}
	userController := controllers.NewUserController(&userService)

	api := r.Group("/api")
	{
		api.POST("/player", middleware.RequireAuth, playerController.CreatePlayer)
		api.POST("/match", middleware.RequireAuth, playerController.AddMatch)
		api.GET("/players", playerController.GetPlayers)
		api.GET("/player/:id", playerController.GetPlayerById)
		api.PATCH("/player/:id", middleware.RequireAuth, playerController.UpdatePlayer)

		api.POST("/toggle_membership/:id", middleware.RequireAuth, playerController.ToggleMembership)
		api.POST("/toggle_payedballs/:id", middleware.RequireAuth, playerController.TogglePayedBalls)
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
