package server

import (
	"encoding/json"
	"net/http"

	"github.com/atla/owndnd/pkg/service"
	log "github.com/sirupsen/logrus"
)

//PartiesHandler is the public item handler interface
type PartiesHandler interface {
	CreateParty(w http.ResponseWriter, r *http.Request)
}

type partiesHandler struct {
	css           service.CharacterSheetsService
	httpResponder HTTPResponder
}

//NewPartiesHandler creates a new item handler
func NewPartiesHandler(css service.CharacterSheetsService, httpResponder HTTPResponder) PartiesHandler {
	return &partiesHandler{
		css:           css,
		httpResponder: httpResponder,
	}
}

// creates a new charactersheet
func (ph *partiesHandler) CreateParty(w http.ResponseWriter, r *http.Request) {

	var createParty CreatePartyDTO
	_ = json.NewDecoder(r.Body).Decode(&createParty)

	log.WithField("charactersheet", createParty).Info("Creating new party")

	/*
		if storedItem, err := csh.css.CreateCharacterSheet(&createParty); err != nil {
			csh.httpResponder.ERROR(w, http.StatusNotFound)
		} else {
			csh.httpResponder.JSON(w, http.StatusOK, storedItem)
		}
	*/
}
