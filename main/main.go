package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nikhileshsirohi/Restaurant-Management/pkg/database"
	"github.com/nikhileshsirohi/Restaurant-Management/pkg/middleware"
	"github.com/nikhileshsirohi/Restaurant-Management/pkg/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9010"
	}
	router := gin.New()

	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
