package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty" form:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" form:"name"`
	Description string             `json:"description" bson:"description" form:"description"`
	Price       float64            `json:"price" bson:"price" form:"price"`
	Category    string             `json:"category" bson:"category" form:"category"`
	Stock       int                `json:"stock" bson:"stock" form:"stock"`
}
