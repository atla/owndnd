package entities

import (
	"time"
)

//CharacterRace type
type CharacterRace int

const (
	crHuman CharacterRace = iota + 1
	crDwarf
	crElve
)

func (cr CharacterRace) String() string {
	return [...]string{"human", "dwarf", "elve"}[cr]
}

//Party data
type Party struct {
	*Entity
	Name string `json:"name"`

	Created    time.Time  `bson:"created,omitempty" json:"created,omitempty"`
	Characters []EntityID `bson:"characters,omitempty" json:"characters,omitempty"`
}
