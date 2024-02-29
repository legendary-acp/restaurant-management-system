package controller

import (
	"github.com/gin-gonic/gin"
	controller "github.com/legendary-acp/restaurant-management-system/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.Login())
}
