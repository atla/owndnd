package service

import (
	"github.com/atla/owndnd/pkg/db"
	"github.com/atla/owndnd/pkg/repository"
)

//Facade ...
type Facade interface {
	CharacterSheetsService() CharacterSheetsService
	PartiesService() PartiesService
	UsersService() UsersService
	RoomsService() RoomsService
}

type facade struct {
	css CharacterSheetsService
	ps  PartiesService
	us  UsersService
	rs  RoomsService
	db  *db.Client
}

//NewFacade creates a new service facade
func NewFacade(db *db.Client) Facade {
	characterSheetRepo := repository.NewMongoDBCharacterSheetRepository(db)
	partiesRepo := repository.NewMongoDBPartiesRepository(db)
	usersRepo := repository.NewMongoDBUsersRepository(db)
	roomsRepo := repository.NewMongoDBRoomsRepository(db)

	return &facade{
		css: NewCharacterSheetsService(characterSheetRepo),
		ps:  NewPartiesService(partiesRepo),
		us:  NewUsersService(usersRepo),
		rs:  NewRoomsService(roomsRepo),
	}
}
func (f *facade) RoomsService() RoomsService {
	return f.rs
}
func (f *facade) CharacterSheetsService() CharacterSheetsService {
	return f.css
}

func (f *facade) PartiesService() PartiesService {
	return f.ps
}
func (f *facade) UsersService() UsersService {
	return f.us
}
