package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// type Account struct {
// 	Name    string `json:"name,omitempty" validate:"required"`
// 	State   string `json:"state,omitempty" validate:"required"`
// 	City    string `json:"city,omitempty" validate:"required"`
// 	Pincode int    `json:"pincode,omitempty" validate:"required"`
// }

type Task struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Date    string             `json:"date,omitempty" validate:"required"`
	Bill    int                `json:"bill,omitempty" validate:"required"`
	Account string             `json:"account,omitempty" validate:"required"`
}
