package controllers

import (
	"context"
	// "fmt"
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



var userCollection *mongo.Collection=database.OpenCollection(database.Client,"user")

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context)  {
		ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	
		
	}

}
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	   var users  []models.User
	  
	   defer cancel()
	   
	   result,err:=userCollection.Find(ctx,bson.M{})
	   if err !=nil{
		c.JSON(http.StatusInternalServerError,dtos.UserResponse{Status: 500,Message: "error",Data: err.Error()})
		return 
	   }
       defer result.Close(ctx)
	   for result.Next(ctx){
		var singleUser models.User
		if err=result.Decode(&singleUser);err!=nil{
			c.JSON(http.StatusInternalServerError,dtos.UserResponse{Status: http.StatusInternalServerError,Message: "error",Data: err.Error()})
		}
		users=append(users, singleUser)
	   }
	   c.JSON(http.StatusOK,dtos.UserResponse{Status: http.StatusOK,Data: users})

	   
	   
	   

	}

}
func SignUp() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx,cancel :=context.WithTimeout(context.Background(),10*time.Second)
	
		var user models.User
		defer cancel()
		// validate the request body
		if err :=c.BindJSON(&user); err!=nil{
			c.JSON(http.StatusBadRequest,dtos.UserResponse{Status: http.StatusBadRequest,Message: "error",Data:  err.Error()})
            return
		}
		newUser:=models.User{
			Id: primitive.NewObjectID(),
			Username: user.Username ,
			Email: user.Email,
			Password: user.Password,

			
					
				}
		result,err:=userCollection.InsertOne(ctx,newUser)
		if err !=nil{
			c.JSON(http.StatusInternalServerError,dtos.UserResponse{Status: http.StatusInternalServerError,Message: "error"})
			return
		}		
		c.JSON(http.StatusCreated,dtos.UserResponse{Status: http.StatusCreated,Message: "user created successfully",Data: result})
			

	}

}
func Login() gin.HandlerFunc {
	
return func(c *gin.Context) {
	
}
}
