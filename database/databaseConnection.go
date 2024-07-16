package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func DBInstance() *mongo.Client{
	Mongodb :="mongodb://localhost:27017"
	fmt.Print(Mongodb)
   client,err:=mongo.NewClient(options.Client().ApplyURI(Mongodb))
   if err!=nil{
	 log.Fatal(err)
   }
   ctx,cancel:=context.WithTimeout(context.Background(),10*time.second)
   defer cancel()
   err=client.Connect(ctx)
   if err!=nil{
	log.Fatal(err)
   }
   fmt.Println("connected to mongodb")
return client
}
var Client *mongo.Client=DBInstance()
func OpenCollection(client *mongo.Client,collectionName string) *mongo.Collection{
	var collection *mongo.Collection=client.Database("restuarant").Collection(collectionName)
return collection
}