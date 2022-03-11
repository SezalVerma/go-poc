package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//---------------- Account Model
type Account struct{
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountType string `json:"type" validate:"required" bson:"type"`
	Balance int `json:"balance" bson:"balance"`
	Currency string `json:"currency" bson:"currency"`
	Aadhar int `json:"aadhar" validate:"required" bson:"aadhar"`
}