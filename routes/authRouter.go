package routes

import (
	controllers "github.com/Rassimdou/Go-Gin_Auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incommingRoutes.POST("user/signup", controllers.Signup())
	incommingRoutes.POST("user/login", controllers.Login())
}
