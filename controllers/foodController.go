package controllers

import (
	"context"

	"net/http"
	"restaurant_mis/database"
	"restaurant_mis/dtos"
	"restaurant_mis/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection=database.OpenCollection(database.Client,"food")

func GetFoods() gin.HandlerFunc{

return func(ctx *gin.Context) {
	var c,cancel=context.WithTimeout(context.Background(),10*time.Second)
    defer cancel()
	var foods []models.Food
	result,_:=foodCollection.Find(c,bson.M{})
	defer result.Close(c)
	for result.Next(c){
		var singleFood models.Food
		if err:=result.Decode(&singleFood);err !=nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":"error while fetching foods"})
		}
		foods=append(foods,singleFood)


      

	}
	ctx.JSON(http.StatusOK,dtos.Response{Status: http.StatusOK,Message: "success",Data: foods})
}
}
func GetFood() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		 var c,cancel=context.WithTimeout(context.Background(),10*time.Second)
		 
		 foodId:=ctx.Param("food_id")
		 var food models.Food
		 defer cancel()

		 err:=foodCollection.FindOne(c,bson.M{"food_id":foodId}).Decode(&food)
		 if err!=nil {
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":"error occured while fetching food "})
			return
		 }
		 ctx.JSON(http.StatusOK,dtos.Response{Status: http.StatusOK,Data: food})

		  
	}

}


func CreateFood() gin.HandlerFunc{
return func(ctx *gin.Context) {
var c,cancel=context.WithTimeout(context.Background(),10*time.Second)
defer cancel()
var food models.Food
if  err:=ctx.BindJSON(&food); err !=nil{
	ctx.JSON(http.StatusBadRequest,gin.H{"error":"error while creating the food"})
	return
}
newFood:=models.Food{
	ID: primitive.NewObjectID(),
	Name: food.Name,
	Menu_id: food.Menu_id,
	Food_image: food.Food_image,
	Price: food.Price,
	Food_id: food.Food_id,

	
}
_,err:=foodCollection.InsertOne(c,newFood)
if err !=nil{
	ctx.JSON(http.StatusInternalServerError,gin.H{"error":"internal server error"})
	return 
}

ctx.JSON(http.StatusCreated,dtos.Response{
Status: http.StatusCreated,
Message: "created successfully",
})

}
}
func UpdateFood() gin.HandlerFunc{
return func(ctx *gin.Context) {}
}
// func HashPassword(password string) string{

// }
// func VerifyPassword(userPassword string,provided string) (bool,string){
	
// }