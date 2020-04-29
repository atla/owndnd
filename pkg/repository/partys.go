package repository

import (
	e "github.com/atla/owndnd/pkg/entities"
)

//PartysRepository repository interface
type PartysRepository interface {
	FindAll() ([]*e.Party, error)
	FindByID(id string) (*e.Party, error)
	FindByName(name string) (*e.Party, error)
	Store(party *e.Party) (*e.Party, error)
	Update(party *e.Party) error
}
