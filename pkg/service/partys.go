package service

import (
	e "github.com/atla/owndnd/pkg/entities"
	r "github.com/atla/owndnd/pkg/repository"
)

//PartysService delives logical functions on top of the charactersheets Repo
type PartysService interface {
	GetPartyByID(id string) (*e.Party, error)
	AddCharacterToParty(party *e.Party, character *e.Character) error
}

type partysService struct {
	repo r.CharacterSheetsRepository
}

//NewPartysService creates a nwe item service
func NewPartysService(characterSheetsRepo r.CharacterSheetsRepository) PartysService {
	return &partysService{
		repo: characterSheetsRepo,
	}
}

func (s *partysService) GetPartyByID(id string) (*e.Party, error) {
	return nil, nil
}

func (s *partysService) AddCharacterToParty(party *e.Party, character *e.Character) error {
	return nil
}
