package api

import (
	"bitbucket.org/arthemismc/core/backend"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	// FactionIn contains informations about a player.
	FactionIn struct {
		// Name is the faction name.
		Name string `json:"name" description:"Faction name." validate:"required,lte=20"`

		// Color is the faction color. It corresponds to the Minecraft wool color.
		// See: https://minecraft.gamepedia.com/Wool
		Color int `json:"color" description:"Faction color. It must be a Minecraft wool color code." validate:"required,min=1,max=15"`
	}

	// FactionUpdateIn contains informations about a player.
	FactionUpdateIn struct {
		// UUID is the player uuid
		UUID string `path:"id" validate:"required"`

		// Name is the faction name.
		Name string `json:"name" description:"Faction name." validate:"required,lte=20"`

		// Color is the faction color. It corresponds to the Minecraft wool color.
		// See: https://minecraft.gamepedia.com/Wool
		Color int `json:"color" description:"Faction color. It must be a Minecraft wool color code." validate:"required,min=1,max=15"`
	}
)

// addFaction adds a new faction to the database.
// It checks if it does not already exist before adding the faction.
func (a *API) addFaction(c *gin.Context, f *FactionIn) (*UUIDOut, error) {
	id, err := a.backend.AddFaction(f.Name, f.Color)
	if err != nil {
		return nil, err
	}

	return &UUIDOut{
		UUID: id,
	}, nil
}

// getFaction returns all informations about a faction corresponding
// to the given ID.
func (a *API) getFaction(c *gin.Context, s *SelectorIn) (*backend.Faction, error) {
	id, err := uuid.Parse(s.UUID)
	if err != nil {
		return nil, err
	}

	faction, err := a.backend.GetFaction(&id)
	if err != nil {
		return nil, err
	}

	return faction, nil
}

// deleteFaction deletes the faction corresponding to the given ID.
func (a *API) deleteFaction(c *gin.Context, s *SelectorIn) (*UUIDOut, error) {
	id, err := uuid.Parse(s.UUID)
	if err != nil {
		return nil, err
	}

	returnedUUID, err := a.backend.DeleteFaction(&id)
	if err != nil {
		return nil, err
	}

	return &UUIDOut{
		UUID: returnedUUID,
	}, nil
}

// updateFaction updates the faction corresponding to the given ID,
// with the new informations.
func (a *API) updateFaction(c *gin.Context, f *FactionUpdateIn) (*UUIDOut, error) {
	id, err := uuid.Parse(f.UUID)
	if err != nil {
		return nil, err
	}

	returnedUUID, err := a.backend.UpdateFaction(&id, f.Name, f.Color)
	if err != nil {
		return nil, err
	}
	return &UUIDOut{
		UUID: returnedUUID,
	}, nil
}
