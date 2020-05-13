package server

import (
	"encoding/json"
	"net/http"

	e "github.com/atla/owndnd/pkg/entities"
	"github.com/atla/owndnd/pkg/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//CharacterSheetHandler is the public item handler interface
type CharacterSheetHandler interface {
	GetCharacterSheets(w http.ResponseWriter, r *http.Request)
	GetCharacterSheetByID(w http.ResponseWriter, r *http.Request)
	DeleteCharacterSheetByID(w http.ResponseWriter, r *http.Request)

	PostCharacterSheet(w http.ResponseWriter, r *http.Request)
	UpdateCharacterSheetByID(w http.ResponseWriter, r *http.Request)
}

type characterSheetHandler struct {
	css           service.CharacterSheetsService
	httpResponder HTTPResponder
}

//NewCharacterSheetHandler creates a new item handler
func NewCharacterSheetHandler(css service.CharacterSheetsService, httpResponder HTTPResponder) CharacterSheetHandler {
	return &characterSheetHandler{
		css:           css,
		httpResponder: httpResponder,
	}
}

// returns the list of item templates
func (csh *characterSheetHandler) GetCharacterSheets(w http.ResponseWriter, r *http.Request) {
	//TODO: check if there is a search/filter
	if characterSheets, err := csh.css.GetCharacterSheets(); err != nil {
		csh.httpResponder.ERROR(w, http.StatusNotFound, err)
	} else {
		csh.httpResponder.JSON(w, http.StatusOK, characterSheets)
	}
}

// returns a single charactersheet
func (csh *characterSheetHandler) GetCharacterSheetByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var id = params["id"]

	if item, err := csh.css.GetCharacterSheetByID(id); err != nil {
		csh.httpResponder.ERROR(w, http.StatusNotFound, err)
	} else {
		csh.httpResponder.JSON(w, http.StatusOK, item)
	}
}

// returns a single charactersheet
func (csh *characterSheetHandler) DeleteCharacterSheetByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var id = params["id"]

	if err := csh.css.DeleteCharacterSheetByID(id); err != nil {
		csh.httpResponder.ERROR(w, http.StatusNotFound, err)
	} else {
		csh.httpResponder.JSON(w, http.StatusOK, "deleted")
	}
}

// creates a new charactersheet
func (csh *characterSheetHandler) UpdateCharacterSheetByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var id = params["id"]
	var charactersheet e.CharacterSheet
	_ = json.NewDecoder(r.Body).Decode(&charactersheet)

	log.WithField("charactersheet", charactersheet.Name).Info("Updating new charactersheet")

	if err := csh.css.UpdateCharacterSheetByID(id, &charactersheet); err != nil {
		csh.httpResponder.ERROR(w, http.StatusNotFound, err)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// creates a new charactersheet
func (csh *characterSheetHandler) PostCharacterSheet(w http.ResponseWriter, r *http.Request) {

	var charactersheet e.CharacterSheet
	_ = json.NewDecoder(r.Body).Decode(&charactersheet)

	log.WithField("charactersheet", charactersheet.Name).Info("Creating new charactersheet")

	if storedItem, err := csh.css.CreateCharacterSheet(&charactersheet); err != nil {
		csh.httpResponder.ERROR(w, http.StatusNotFound, err)
	} else {
		csh.httpResponder.JSON(w, http.StatusOK, storedItem)
	}
}
