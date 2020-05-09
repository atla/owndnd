package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EntityID ...
type EntityID string

// NewEntity creates a new base entity
func NewEntity() *Entity {
	return &Entity{
		ID: EntityID(primitive.NewObjectID().Hex()),
	}
}

//Entity ...
type Entity struct {
	ID EntityID `bson:"_id,omitempty" json:"id,omitempty"`
}
