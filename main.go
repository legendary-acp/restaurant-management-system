package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/legendary-acp/restaurant-management-system/database"
	"github.com/legendary-acp/restaurant-management-system/middleware"
	routes "github.com/legendary-acp/restaurant-management-system/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access Granted"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access Granted for 2"})
	})

	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	router.Use(middleware.Authenticate())
	routes.UserRoutes(router)

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)

	router.Run(":" + port)
}
