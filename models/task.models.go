package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Name    string `json:"name,omitempty" validate:"required"`
	State   string `json:"state,omitempty" validate:"required"`
	City    string `json:"city,omitempty" validate:"required"`
	Pincode string `json:"pincode,omitempty" validate:"required"`
}

type Power struct {
	Name    string  `json:"name,omitempty" validate:"required"`
	Bill    int     `json:"bill,omitempty" validate:"required"`
	Used    int     `json:"used,omitempty" validate:"required"`
	Date    string  `json:"date,omitempty" validate:"required"`
	Account Account `json:"account,omitempty" validate:"required"`
}

type Gas struct {
	Name    string  `json:"name,omitempty" validate:"required"`
	Bill    int     `json:"bill,omitempty" validate:"required"`
	Used    int     `json:"used,omitempty" validate:"required"`
	Date    string  `json:"date,omitempty" validate:"required"`
	Account Account `json:"account,omitempty" validate:"required"`
}

type Task struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Bill  int                `json:"bill,omitempty" validate:"required"`
	Date  string             `json:"date,omitempty" validate:"required"`
	Gas   Gas                `json:"gas,omitempty" validate:"required"`
	Power Power              `json:"power,omitempty" validate:"required"`
}
