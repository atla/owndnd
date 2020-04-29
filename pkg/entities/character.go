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

// NewCharacterFrom creates a new character from charactersheet
func NewCharacterFrom(characterSheet *CharacterSheet) *Character {
	return &Character{
		Entity: &Entity{
			ID: characterSheet.ID,
		},
		Name:        characterSheet.Name,
		Description: characterSheet.Description,
		Race:        characterSheet.Race,
	}
}
