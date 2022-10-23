package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type City struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	Latitude  float64            `json:"latitude,omitempty" validate:"required"`
	Longitude float64            `json:"longitude,omitempty" validate:"required"`
}
