package repository

import (
	"github.com/atla/owndnd/pkg/db"
	e "github.com/atla/owndnd/pkg/entities"
)

//PartysRepository repository interface
type PartiesRepository interface {
	//FindAll() ([]*e.Party, error)
	//	FindByID(id string) (*e.Party, error)
	//FindByName(name string) (*e.Party, error)
	Store(party *e.Party) (*e.Party, error)
	//Update(party *e.Party) error
}

//--- Implementations

type partiesRepo struct {
	*GenericRepo
}

//NewMongoDBPartiesRepository creates a new mongodb partiesrep
func NewMongoDBPartiesRepository(db *db.Client) PartiesRepository {
	return &partiesRepo{
		GenericRepo: &GenericRepo{
			db:         db,
			collection: "parties",
			generator: func() interface{} {
				return &e.Party{}
			},
		},
	}
}

func (pr *partiesRepo) Store(party *e.Party) (*e.Party, error) {
	result, err := pr.GenericRepo.Store(party)
	return result.(*e.Party), err
}
