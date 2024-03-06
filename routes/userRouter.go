package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/legendary-acp/restaurant-management-system/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
