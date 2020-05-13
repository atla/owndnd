package entities

//Character header data is an abbrevated version of the character sheet
type Character struct {
	*Entity
	Name        string `json:"name"`
	Description string `json:"description"`
	Race        Race   `json:"race"`
	Class       Class  `json:"class"`
}

// NewCharacter creates a new character
func NewCharacter() *Character {
	return &Character{}
}
