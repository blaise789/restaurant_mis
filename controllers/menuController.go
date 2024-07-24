package controllers

import (
	"context"
	"fmt"
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

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		var menus []models.Menu

		defer cancel()
		result, err := menuCollection.Find(c, bson.M{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "failed to fetch", Data: err.Error()})
			return
		}
		defer result.Close(c)
		for result.Next(c) {
			var menu models.Menu
			if result.Decode(&menu); err != nil {
				ctx.JSON(http.StatusInternalServerError, dtos.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
				return
			}
			menus = append(menus, menu)
		}
		ctx.JSON(http.StatusOK, dtos.Response{Status: http.StatusOK, Message: "success", Data: menus})

	}

}

func GetMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()
		menuId := ctx.Param("menu_id")
	    fmt.Println(menuId)

		var menu models.Menu
		objectId, _ := primitive.ObjectIDFromHex(menuId)

		err:= menuCollection.FindOne(c, bson.M{"_id":objectId}).Decode(&menu) 
	

		if err!=nil {
			ctx.JSON(http.StatusNotFound,gin.H{"message":"menu not found"})
			return
		}
		
		ctx.JSON(http.StatusOK,dtos.Response{Status: http.StatusOK,Message: "success",Data: menu})
		

	}
}
func CreateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var menu models.Menu
		err := ctx.BindJSON(&menu)
		defer cancel()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, dtos.Response{Status: http.StatusBadRequest, Message: "bad request"})
			return
		}
		newMenu := models.Menu{
			Id:         primitive.NewObjectID(),
			Name:       menu.Name,
			Category:   menu.Category,
			Start_date: menu.Start_date,
			End_date:   menu.End_date,
		}
		_, e := menuCollection.InsertOne(c, newMenu)
		if e != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": e.Error()})
			return

		}
		ctx.JSON(http.StatusCreated, dtos.Response{Status: http.StatusCreated, Message: "successfully created"})

	}
}
func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
