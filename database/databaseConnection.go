package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func DBInstance() *mongo.Client{
	mongoUri :="mongodb://localhost:27017"
	clientOptions:=options.Client().ApplyURI(mongoUri)
   client,err:=mongo.Connect(context.TODO(),clientOptions )
   if err!=nil{
	 log.Fatal(err)
   }
   err = client.Ping(context.TODO(), nil)
   if err!=nil{
	log.Fatal(err)
   }
   log.Println("connected to mongodb")
return client
}
var Client *mongo.Client=DBInstance()
func OpenCollection(client *mongo.Client,collectionName string) *mongo.Collection{
	var collection *mongo.Collection=client.Database("restaurant_mis").Collection(collectionName)
return collection
}