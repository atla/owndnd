package server

import (
	"encoding/json"
	"net/http"

	"github.com/atla/owndnd/pkg/entities"
	"github.com/atla/owndnd/pkg/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//PartiesHandler is the public item handler interface
type PartiesHandler interface {
	CreateParty(w http.ResponseWriter, r *http.Request)
	UpdateParty(w http.ResponseWriter, r *http.Request)
	DeleteParty(w http.ResponseWriter, r *http.Request)
}

type partiesHandler struct {
	service       service.PartiesService
	httpResponder HTTPResponder
}

//NewPartiesHandler creates a new item handler
func NewPartiesHandler(srv service.PartiesService, httpResponder HTTPResponder) PartiesHandler {
	return &partiesHandler{
		service:       srv,
		httpResponder: httpResponder,
	}
}

// creates a new charactersheet
func (handler *partiesHandler) CreateParty(w http.ResponseWriter, r *http.Request) {

	var createParty service.CreatePartyDTO
	_ = json.NewDecoder(r.Body).Decode(&createParty)

	log.WithField("charactersheet", createParty).Info("Creating new party")

	if storedItem, err := handler.service.CreateParty(&createParty); err != nil {
		handler.httpResponder.ERROR(w, http.StatusNotFound)
	} else {
		handler.httpResponder.JSON(w, http.StatusOK, storedItem)
	}
}

// creates a new charactersheet
func (handler *partiesHandler) UpdateParty(w http.ResponseWriter, r *http.Request) {

	var party entities.Party
	_ = json.NewDecoder(r.Body).Decode(&party)

	params := mux.Vars(r)
	var id = params["id"]

	party.ID = entities.EntityID(id)

	log.WithField("party", party).Info("Updating new party")

	if err := handler.service.UpdateParty(id, &party); err != nil {
		handler.httpResponder.ERROR(w, http.StatusInternalServerError)
	} else {
		handler.httpResponder.JSON(w, http.StatusOK, nil)
	}

}

// creates a new charactersheet
func (handler *partiesHandler) DeleteParty(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var id = params["id"]

	log.WithField("party id", id).Info("deleting party")

	if err := handler.service.DeletePartyByID(id); err != nil {
		handler.httpResponder.ERROR(w, http.StatusInternalServerError)
	} else {
		handler.httpResponder.JSON(w, http.StatusOK, nil)
	}

}
