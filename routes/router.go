package routes

import (
	"os"
	"strings"
	"time"

	"github.com/cesart18/ranking_app/controllers"
	"github.com/cesart18/ranking_app/middleware"
	"github.com/cesart18/ranking_app/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	origins := strings.Split(allowedOrigins, ",")
	config := cors.Config{
		AllowOrigins:     origins,
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
		api.GET("/players", playerController.GetPlayers)
		api.GET("/player/:id", playerController.GetPlayerById)
		api.PATCH("/player/:id", middleware.RequireAuth, playerController.UpdatePlayer)
		api.POST("/double_point/:id", middleware.RequireAuth, playerController.AddDoublePoint)
		api.POST("/single_point/:id", middleware.RequireAuth, playerController.AddSinglePoint)

		api.POST("/toggle_membership/:id", middleware.RequireAuth, playerController.ToggleMembership)
		api.POST("/toggle_payedballs/:id", middleware.RequireAuth, playerController.TogglePayedBalls)
		api.DELETE("/player/:id", middleware.RequireAuth, playerController.DeletePlayer)
	}
	auth := r.Group("/auth")
	{
		auth.POST("/signup", userController.Signup)
		auth.POST("/login", userController.Login)
		auth.POST("/logout", userController.Logout)
		auth.GET("/validate", middleware.RequireAuth, userController.Validate)
	}

	return r
}
