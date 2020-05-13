package entities

import (
	"time"
)

//Race type
type Race int

const (
	rHuman Race = iota + 1
	rDwarf
	rElve
)

func (cr Race) String() string {
	return [...]string{"human", "dwarf", "elve"}[cr]
}

//Class type
type Class int

const (
	cWarrior Class = iota + 1
	cWizard
	cRanger
)

func (cr Class) String() string {
	return [...]string{"warrior", "wzard", "ranger"}[cr]
}

// Attribute data
type Attribute struct {
	Name  string `json:"name"`
	Value int32  `json:"value"`
}

//CharacterSheet data
type CharacterSheet struct {
	*Entity     `bson:",inline"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Race        Race   `json:"race"`
	Class       Class  `json:"class"`

	CurrentHitPoints int32 `json:"currentHitPoints"`
	MaxHitPoints     int32 `json:"maxHitPoints"`
	ArmorClass       int32 `json:"armorClass"`

	Created    time.Time    `bson:"created" json:"created,omitempty"`
	Attributes []*Attribute `bson:"attributes" json:"attributes,omitempty"`

	PersonalityTraits string `json:"personalityTraits,omitempty"`
}

// AsCharacter creates a new character from charactersheet
func (cs *CharacterSheet) AsCharacter() *Character {
	return &Character{

		Name:        cs.Name,
		Description: cs.Description,
		Race:        cs.Race,
	}
}
