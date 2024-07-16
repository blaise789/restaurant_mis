package routes
import (
	controller "restaurant_mis/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine){
incomingRoutes.GET("/orderItems",controller.GetOrderItems())
incomingRoutes.GET("/orderItems/:order_item_id",controller.GetOrderItem())
incomingRoutes.GET("/orderItems-order/:order_id",controller.GetOrderItemsByOrder())
incomingRoutes.POST("/orderItems",controller.CreateOrderItem())
incomingRoutes.PATCH("/orderItems/:order_item_id",controller.UpdateOrderItem())

}