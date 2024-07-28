package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	Id primitive.ObjectID `bson:"_id"`
	Name  string `json:"name"`
	Category string `json:"category"`
	Start_date  *time.Time `json:"start_date,omitempty"`
	End_date  *time.Time `json:"end_date,omitempty"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}