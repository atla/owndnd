package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HasID ...
type HasID interface {
	SetID(id primitive.ObjectID)
	GetID() primitive.ObjectID
}

// EntityID ...
type EntityID primitive.ObjectID

//Entity ...
type Entity struct {
	ID EntityID `bson:"_id,omitempty" json:"id,omitempty"`
}

// SetID ...
func (e *Entity) SetID(id EntityID) {
	e.ID = id
}

// GetID ...
func (e *Entity) GetID() EntityID {
	return e.ID
}
