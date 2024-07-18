package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct{
	ID primitive.ObjectID  `bson:"_id"`
	Name  *string      `json:"name," validate:"required,min=3,max=100"`
	Price  *float64     `json:"price" validate:"required,max=10"`
	Food_image   * string  `json:"food_image,omitempty"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	Food_id       string `json:"food_id,omitempty"`
	Menu_id *string   `json:"menu_id,omitempty" validate:"required"` 
}