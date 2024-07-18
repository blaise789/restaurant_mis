package main

import (
	"os"
	// "restaurant_mis/models"
	
	"restaurant_mis/routes"

	"github.com/gin-gonic/gin"
	// "restaurant_mis/middlewares"
)
func main(){
 port:= os.Getenv("PORT")
 if port ==""{
	port ="8500"
 }
 router :=gin.New();
 router.GET("/",func(ctx *gin.Context) {
	ctx.JSON(200,gin.H {
		"message":"hello",
	})
 },

)
 router.Use(gin.Logger())
 routes.UserRoutes(router)
//  intercept routes with auth
//  router.Use(middlewares.Authentication())
 routes.FoodRoutes(router)
 routes.MenuRoutes(router)
 routes.TableRoutes(router)
 routes.OrderRoutes(router)
 routes.OrderItemRoutes(router)
 routes.InvoiceRoutes(router)
 router.Run(":"+port)

 
 
 
 
 
}