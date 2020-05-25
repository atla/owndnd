package handler

import (
	"net/http"

	e "github.com/atla/owndnd/pkg/entities"
	"github.com/atla/owndnd/pkg/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//CharactersHandler ...
type CharactersHandler struct {
	Service service.CharacterSheetsService
}

//GetCharacters returns the list of item templates
func (csh *CharactersHandler) GetCharacters(c *gin.Context) {

	if characters, err := csh.Service.GetCharacterSheets(); err == nil {
		c.JSON(http.StatusOK, characters)
	} else {
		c.Error(err)
	}
}

//GetCharacterByID returns a single charactersheet
func (csh *CharactersHandler) GetCharacterByID(c *gin.Context) {

	id := c.Query("id")

	if character, err := csh.Service.GetCharacterSheetByID(id); err == nil {
		c.JSON(http.StatusOK, character)
	} else {
		c.Error(err)
	}
}

//DeleteCharacterByID returns a single charactersheet
func (csh *CharactersHandler) DeleteCharacterByID(c *gin.Context) {

	id := c.Query("id")

	if err := csh.Service.DeleteCharacterSheetByID(id); err == nil {
		c.JSON(http.StatusOK, "deleted")
	} else {
		c.Error(err)
	}
}

//UpdateCharacterByID creates a new charactersheet
func (csh *CharactersHandler) UpdateCharacterByID(c *gin.Context) {

	id := c.Query("id")
	var character e.CharacterSheet
	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.WithField("character", character.Name).Info("Updating character")

	if err := csh.Service.UpdateCharacterSheetByID(id, &character); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": "updated character"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//PostCharacter ... creates a new charactersheet
func (csh *CharactersHandler) PostCharacter(c *gin.Context) {

	var character e.CharacterSheet
	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.WithField("character", character.Name).Info("Creating new character")

	if newCharacter, err := csh.Service.CreateCharacterSheet(&character); err == nil {
		c.JSON(http.StatusOK, newCharacter)
	} else {
		c.Error(err)
	}
}
