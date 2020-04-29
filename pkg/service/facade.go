package service

import (
	"github.com/atla/owndnd/pkg/db"
	"github.com/atla/owndnd/pkg/repository"
)

//Facade ...
type Facade interface {
	CharacterSheetsService() CharacterSheetsService
	//	PartysService() PartysService
}

type facade struct {
	css CharacterSheetsService
	//ps  PartysService

	db *db.Client
}

//NewFacade creates a new service facade
func NewFacade(db *db.Client) Facade {

	characterSheetRepo := repository.NewMongoDBCharacterSheetRepository(db)

	return &facade{
		css: NewCharacterSheetsService(characterSheetRepo),
	}
}

func (f *facade) CharacterSheetsService() CharacterSheetsService {
	return f.css
}
