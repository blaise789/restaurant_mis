package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id primitive.ObjectID `bson:"_id"`
	Username string `json:"username" `
	Password string `json:"password" validate:"required"`
	Email string `json:"email" validate:"required" `
}