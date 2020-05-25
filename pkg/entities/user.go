package entities

import "time"

//User connects to login credentials to a user object
type User struct {
	*Entity `bson:",inline"`

	// main *unique* reference id, its additional to ID because it might be a list of reference ids in the future so a user
	// can merge multiple accounts (refids) into the same user object
	RefID string `json:"refid,omitempty"`

	Name     string    `json:"name"`
	Nickname string    `json:"nickname"`
	Email    string    `json:"email"`
	Created  time.Time `bson:"created" json:"created,omitempty"`
	LastSeen time.Time `bson:"lastSeen" json:"lastSeen,omitempty"`
	Picture  string    `json:"picture"`

	// is set to false after the first PUT request
	IsNewUser bool `json:"isNewUser"`
}

// NewUser creates a new user
func NewUser() *User {
	return &User{}
}
