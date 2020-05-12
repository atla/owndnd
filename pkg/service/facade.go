package service

import (
	"github.com/atla/owndnd/pkg/db"
	"github.com/atla/owndnd/pkg/repository"
)

//Facade ...
type Facade interface {
	CharacterSheetsService() CharacterSheetsService
	PartiesService() PartiesService
}

type facade struct {
	css CharacterSheetsService
	ps  PartiesService

	db *db.Client
}

//NewFacade creates a new service facade
func NewFacade(db *db.Client) Facade {

	characterSheetRepo := repository.NewMongoDBCharacterSheetRepository(db)
	partiesRepo := repository.NewMongoDBPartiesRepository(db)

	return &facade{
		css: NewCharacterSheetsService(characterSheetRepo),
		ps:  NewPartiesService(partiesRepo),
	}
}

func (f *facade) CharacterSheetsService() CharacterSheetsService {
	return f.css
}

func (f *facade) PartiesService() PartiesService {
	return f.ps
}
