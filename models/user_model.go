package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// -----------------   User Model
type User struct{
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" validate:"required" bson:"name"`
	Gender string `json:"gender,omitempty" bson:"gender,omitempty"`
	Address string `json:"address,omitempty" validate:"required" bson:"address,omitempty"`
	Phone int `json:"phone" validate:"required" bson:"phone"`
	Aadhar int `json:"aadhar" validate:"required" bson:"aadhar"`
	PanNumber string `json:"pan" bson:"pan"`
}