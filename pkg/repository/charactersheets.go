package repository

import (
	"github.com/atla/owndnd/pkg/db"
	e "github.com/atla/owndnd/pkg/entities"
)

//--- Interface Definitions

//CharacterSheetsRepository repository interface
type CharacterSheetsRepository interface {
	FindAll() ([]*e.CharacterSheet, error)
	FindByID(id string) (*e.CharacterSheet, error)
	FindByName(name string) ([]*e.CharacterSheet, error)
	Store(characterSheet *e.CharacterSheet) (*e.CharacterSheet, error)
	Update(characterSheet *e.CharacterSheet) error
	Delete(id e.EntityID) error
}

//--- Implementations

type characterSheetRepository struct {
	*GenericRepo
}

//NewMongoDBCharacterSheetRepository creates a new mongodb characterSheetRepository
func NewMongoDBCharacterSheetRepository(db *db.Client) CharacterSheetsRepository {
	return &characterSheetRepository{
		GenericRepo: &GenericRepo{
			db:         db,
			collection: "charactersheets",
			generator: func() interface{} {
				return &e.CharacterSheet{}
			},
		},
	}
}

func (ir *characterSheetRepository) FindByID(id string) (*e.CharacterSheet, error) {
	result, err := ir.GenericRepo.FindByID(id)
	if err == nil {
		return result.(*e.CharacterSheet), nil
	}
	return nil, err
}

func (ir *characterSheetRepository) FindByName(name string) ([]*e.CharacterSheet, error) {
	var results []*e.CharacterSheet

	if err := ir.GenericRepo.FindAllWithParam(
		db.NewQueryParams(db.QueryParam{Key: "name", Value: name}),
		func(elem interface{}) {
			results = append(results, elem.(*e.CharacterSheet))
		}); err != nil {
		return nil, err
	}
	return results, nil
}

func (ir *characterSheetRepository) FindAll() ([]*e.CharacterSheet, error) {
	var results []*e.CharacterSheet
	if err := ir.GenericRepo.FindAll(func(elem interface{}) {
		results = append(results, elem.(*e.CharacterSheet))
	}); err != nil {
		return nil, err
	}
	return results, nil
}

func (ir *characterSheetRepository) Update(charachterSheet *e.CharacterSheet) error {
	return ir.GenericRepo.Update(charachterSheet, charachterSheet.ID)
}

func (ir *characterSheetRepository) Delete(id e.EntityID) error {
	return ir.GenericRepo.Delete(id)
}

func (ir *characterSheetRepository) Store(characterSheet *e.CharacterSheet) (*e.CharacterSheet, error) {
	result, err := ir.GenericRepo.Store(characterSheet)
	return result.(*e.CharacterSheet), err
}
