package routes
import (
	"github.com/gin-gonic/gin"
	controller "restaurant_mis/controllers"
)
func UserRoutes(incomingRoutes *gin.Engine){
	// handing routes
	incomingRoutes.GET("/users",controller.GetUsers())
	incomingRoutes.GET("/users/:id",controller.GetUser())
	incomingRoutes.POST("/users/signup",controller.SignUp())
	incomingRoutes.POST("/users/login",controller.Login())
}
