package handler

import (
	"net/http"

	"github.com/atla/owndnd/pkg/entities/rooms"
	"github.com/atla/owndnd/pkg/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//RoomsHandler ...
type RoomsHandler struct {
	Service service.RoomsService
}

//GetRooms returns the list of item templates
func (handler *RoomsHandler) GetRooms(c *gin.Context) {

	if rooms, err := handler.Service.FindAll(); err == nil {
		c.JSON(http.StatusOK, rooms)
	} else {
		c.Error(err)
	}
}

//PostRoom ... creates a new charactersheet
func (handler *RoomsHandler) PostRoom(c *gin.Context) {

	var room rooms.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.WithField("room", room.Name).Info("Creating new room")

	if room, err := handler.Service.Store(&room); err == nil {
		c.JSON(http.StatusOK, room)
	} else {
		c.Error(err)
	}
}

//PutRoom ... Updates a room
func (handler *RoomsHandler) PutRoom(c *gin.Context) {

	id := c.Query("id")
	var room rooms.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.WithField("room", room.Name).Info("Updating room")

	if err := handler.Service.Update(id, &room); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": "updated room"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

/*
//GetCharacterByID returns a single charactersheet
func (handler *RoomsHandler) GetCharacterByID(c *gin.Context) {

	id := c.Query("id")

	if character, err := handler.Service.GetCharacterSheetByID(id); err == nil {
		c.JSON(http.StatusOK, character)
	} else {
		c.Error(err)
	}
}

//DeleteCharacterByID returns a single charactersheet
func (handler *RoomsHandler) DeleteCharacterByID(c *gin.Context) {

	id := c.Query("id")

	if err := handler.Service.DeleteCharacterSheetByID(id); err == nil {
		c.JSON(http.StatusOK, "deleted")
	} else {
		c.Error(err)
	}
}

//UpdateCharacterByID creates a new charactersheet
func (handler *RoomsHandler) UpdateCharacterByID(c *gin.Context) {

	id := c.Query("id")
	var character e.CharacterSheet
	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.WithField("character", character.Name).Info("Updating character")

	if err := handler.Service.UpdateCharacterSheetByID(id, &character); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": "updated character"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//PostCharacter ... creates a new charactersheet
func (handler *RoomsHandler) PostCharacter(c *gin.Context) {

	var character e.CharacterSheet
	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.WithField("character", character.Name).Info("Creating new character")

	if newCharacter, err := handler.Service.CreateCharacterSheet(&character); err == nil {
		c.JSON(http.StatusOK, newCharacter)
	} else {
		c.Error(err)
	}
}
*/
