package service

import (
	e "github.com/atla/owndnd/pkg/entities"
	r "github.com/atla/owndnd/pkg/repository"
)

//--- Interface Definitions

//CharacterSheetsService delives logical functions on top of the charactersheets Repo
type CharacterSheetsService interface {
	GetCharacterSheets() ([]*e.CharacterSheet, error)
	GetCharacterSheetByID(id string) (*e.CharacterSheet, error)
	CreateCharacterSheet(characterSheet *e.CharacterSheet) (*e.CharacterSheet, error)
}

//--- Implementations

type characterSheetsService struct {
	repo r.CharacterSheetsRepository
}

//NewCharacterSheetsService creates a nwe item service
func NewCharacterSheetsService(characterSheetsRepo r.CharacterSheetsRepository) CharacterSheetsService {
	return &characterSheetsService{
		repo: characterSheetsRepo,
	}
}

func (s *characterSheetsService) GetCharacterSheets() ([]*e.CharacterSheet, error) {
	return s.repo.FindAll()
}

func (s *characterSheetsService) GetCharacterSheetByID(id string) (*e.CharacterSheet, error) {
	return s.repo.FindByID(id)
}

func (s *characterSheetsService) CreateCharacterSheet(characterSheet *e.CharacterSheet) (*e.CharacterSheet, error) {

	characterSheet.Entity = e.NewEntity()

	return s.repo.Store(characterSheet)

}
