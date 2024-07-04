package main
import (
	"os"
	// "restaurant_mis/models"
	// "restaurant_mis/controllers"
	"github.com/gin-gonic/gin"
	// "restaurant_mis/routes"
	// "restaurant_mis/middlewares"
	// "restaurant_mis/database"
	// "go.mongodb.org/mongo-driver/mongo"
	
	
)
// var foodCollection *mongo
func main(){
 port:= os.Getenv("PORT")
 if port ==""{
	port ="8500"
 }
 router :=gin.Default()
 router.GET("/",func(ctx *gin.Context) {
	ctx.JSON(200,gin.H {
		"message":"hello",
	})
 },

)
//  router.Use(gin.Logger())
//  routes.UserRoutes(router)
//  router.Use(middlewares.Authentication())
 
//  routes.FoodRoutes(router)
//  routes.MenuRoutes(router)
//  routes.TableRoutes(router)
//  routes.OrderRoutes(router)
//  routes.OrderItemRoutes(router)
//  routes.InvoiceRoutes(router)
 router.Run(":"+port)

 
 
 
 
 
}