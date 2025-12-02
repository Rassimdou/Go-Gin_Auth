package routes

import (
	controllers "github.com/Rassimdou/Go-Gin_Auth/controllers"
	"github.com/Rassimdou/Go-Gin_Auth/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/user/:userId", controllers.GetUser())
}
