package routes

import (
	"khvnkay/restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/ordersItems", controllers.GetOrderItems())
	incomingRoutes.GET("/ordersItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/ordersItems-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/ordersItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/ordersItems/:orderItem_id", controllers.UpdateOrderItem())
}
