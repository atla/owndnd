package repository

import (
	"github.com/atla/owndnd/pkg/db"
	e "github.com/atla/owndnd/pkg/entities"
)

//PartiesRepository repository interface
type PartiesRepository interface {
	//FindAll() ([]*e.Party, error)
	FindByID(id string) (*e.Party, error)
	//FindByName(name string) (*e.Party, error)
	Store(party *e.Party) (*e.Party, error)
	Update(id string, party *e.Party) error
	Delete(id string) error
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

func (pr *partiesRepo) FindByID(id string) (*e.Party, error) {
	result, err := pr.GenericRepo.FindByID(id)
	return result.(*e.Party), err
}

func (pr *partiesRepo) Update(id string, party *e.Party) error {
	return pr.GenericRepo.Update(party, e.EntityID(id))
}

func (pr *partiesRepo) Delete(id string) error {
	return pr.GenericRepo.Delete(e.EntityID(id))
}
