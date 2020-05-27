package repository

import (
	"github.com/atla/owndnd/pkg/db"
	e "github.com/atla/owndnd/pkg/entities"
	r "github.com/atla/owndnd/pkg/entities/rooms"
)

//--- Interface Definitions

//RoomsRepository repository interface
type RoomsRepository interface {
	FindAll() ([]*r.Room, error)
	FindByID(id string) (*r.Room, error)
	FindByName(name string) ([]*r.Room, error)
	Store(room *r.Room) (*r.Room, error)
	Update(id string, room *r.Room) error
	Delete(id string) error
}

//--- Implementations

type roomsRepository struct {
	*GenericRepo
}

//NewMongoDBRoomsRepository creates a new mongodb repoRepository
func NewMongoDBRoomsRepository(db *db.Client) RoomsRepository {
	return &roomsRepository{
		GenericRepo: &GenericRepo{
			db:         db,
			collection: "rooms",
			generator: func() interface{} {
				return &r.Room{}
			},
		},
	}
}

func (repo *roomsRepository) FindByID(id string) (*r.Room, error) {
	result, err := repo.GenericRepo.FindByID(id)
	if err == nil {
		return result.(*r.Room), nil
	}
	return nil, err
}

func (repo *roomsRepository) FindByName(name string) ([]*r.Room, error) {
	results := make([]*r.Room, 0)

	if err := repo.GenericRepo.FindAllWithParam(
		db.NewQueryParams(db.QueryParam{Key: "name", Value: name}),
		func(elem interface{}) {
			results = append(results, elem.(*r.Room))
		}); err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *roomsRepository) FindAll() ([]*r.Room, error) {
	results := make([]*r.Room, 0)
	if err := repo.GenericRepo.FindAll(func(elem interface{}) {
		results = append(results, elem.(*r.Room))
	}); err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *roomsRepository) Update(id string, charachterSheet *r.Room) error {
	return repo.GenericRepo.Update(charachterSheet, id)
}

func (repo *roomsRepository) Delete(id string) error {
	return repo.GenericRepo.Delete(id)
}

func (repo *roomsRepository) Store(rep *r.Room) (*r.Room, error) {

	rep.Entity = e.NewEntity()

	result, err := repo.GenericRepo.Store(rep)
	return result.(*r.Room), err
}
