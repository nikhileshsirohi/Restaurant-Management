package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhileshsirohi/Restaurant-Management/pkg/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFood())
	incomingRoutes.POST("/foods", controllers.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controllers.UpdateFood())
	//PUT is used for updating or replacing a resource entirely
	//PATCH is used for making partial updates to a resource
}
