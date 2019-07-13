package api

import (
	"bitbucket.org/arthemismc/core/backend"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	// SelectorIn represents the selector used to target a player.
	SelectorIn struct {
		// UUID is the player uuid
		UUID string `path:"id" validate:"required"`
	}

	// PlayerIn contains informations about a player.
	PlayerIn struct {
		// Nickname is the player nickname
		Nickname string `json:"nickname" description:"Player nickname." validate:"required"`
		// Register is true if the player is registered
		Register bool `json:"register" description:"True if the player is registered." default:"false"`
	}

	// RegisterIn is a struct used in to registered (or unregistered) an existing player.
	RegisterIn struct {
		// UUID is the player uuid.
		UUID string `path:"id" description:"Player uuid." validate:"required"`
		// Register is true if the player is registered
		Register bool `query:"register" description:"True if the player is registered." default:"true"`
	}

	// PlayerFactionIn is a struct used to edit the faction of a player.
	PlayerFactionIn struct {
		// UUID is the player uuid.
		UUID string `path:"id" description:"Player uuid." validate:"required"`
		// Faction is the player faction.
		Faction string `json:"faction" description:"Player faction." validate:"required"`
	}

	// UUIDOut is a struct used to return a uuid.
	UUIDOut struct {
		// UUID returned
		UUID *uuid.UUID `json:"uuid" description:"Edited player uuid."`
	}
)

// addPlayer adds a new player to the database.
// It checks if it does not already exist before adding the player.
func (a *API) addPlayer(c *gin.Context, p *PlayerIn) (*UUIDOut, error) {
	// inserts new player in the database
	id, err := a.backend.AddPlayer(p.Nickname, p.Register)
	if err != nil {
		return nil, err
	}

	return &UUIDOut{
		UUID: id,
	}, nil
}

// getPlayer returns all informations about a player corresponding
// to the given ID.
func (a *API) getPlayer(c *gin.Context, s *SelectorIn) (*backend.Player, error) {
	id, err := uuid.Parse(s.UUID)
	if err != nil {
		return nil, err
	}

	player, err := a.backend.GetPlayer(&id)
	if err != nil {
		return nil, err
	}

	return player, nil
}

// deletePlayer deletes the player corresponding to the given ID.
func (a *API) deletePlayer(c *gin.Context, s *SelectorIn) (*UUIDOut, error) {
	id, err := uuid.Parse(s.UUID)
	if err != nil {
		return nil, err
	}

	returnedUUID, err := a.backend.DeletePlayer(&id)
	if err != nil {
		return nil, err
	}

	return &UUIDOut{
		UUID: returnedUUID,
	}, nil
}

// registerPlayer edits the "register" status of the player corresponding to the given UUID.
func (a *API) registerPlayer(c *gin.Context, r *RegisterIn) (*UUIDOut, error) {
	id, err := uuid.Parse(r.UUID)
	if err != nil {
		return nil, err
	}

	returnedUUID, err := a.backend.RegisterPlayer(&id, r.Register)
	if err != nil {
		return nil, err
	}

	return &UUIDOut{
		UUID: returnedUUID,
	}, nil
}

// setPlayerFaction sets faction to a player.
func (a *API) setPlayerFaction(c *gin.Context, p *PlayerFactionIn) (*UUIDOut, error) {
	// parses player uuid
	id, err := uuid.Parse(p.UUID)
	if err != nil {
		return nil, err
	}

	// parses faction uuid
	factionID, err := uuid.Parse(p.Faction)
	if err != nil {
		return nil, err
	}

	returnedUUID, err := a.backend.SetPlayerFaction(&id, &factionID)
	if err != nil {
		return nil, err
	}

	return &UUIDOut{
		UUID: returnedUUID,
	}, nil
}
