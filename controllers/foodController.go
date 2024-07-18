package controllers

import (
	"context"
	"net/http"
	"restaurant_mis/database"
	"restaurant_mis/models"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection=database.OpenCollection(database.Client,"food")

func GetFoods() gin.HandlerFunc{

return func(ctx *gin.Context) {}
}
func GetFood() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		 var c,cancel=context.WithTimeout(context.Background(),100*time.Second)
		 fmt.Println(c)
		 foodId:=ctx.Param("food_id")
		 var food models.Food
		 err:=foodCollection.FindOne(ctx,bson.M{"food_id":foodId}).Decode(&food)
		 if err!=nil {
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":"error occured while fetching food "})
		 }
		 defer cancel()
		 ctx.JSON(http.StatusOK,food)

		  
	}

}


func CreateFood() gin.HandlerFunc{
return func(ctx *gin.Context) {}
}
func UpdateFood() gin.HandlerFunc{
return func(ctx *gin.Context) {}
}
// func HashPassword(password string) string{

// }
// func VerifyPassword(userPassword string,provided string) (bool,string){
	
// }