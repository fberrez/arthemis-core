package backend

import (
	"github.com/google/uuid"
)

type Player struct {
	UUID       uuid.UUID `json:"uuid" description:"Player UUID."`
	Nickname   string    `json:"nickname"description:"Player nickname."`
	Registered bool      `json:"registered" description:"True if the player is registered."`
	Faction    uuid.UUID `json:"faction" description:"Player faction."`
}

// AddPlayer adds a new player to the database.
// It returns the uuid of the player added.
func (b *Backend) AddPlayer(nickname string, registered bool) (*uuid.UUID, error) {
	// prepares the query
	stmt, err := b.db.Prepare("INSERT INTO player(nickname, registered) VALUES ($1, $2) RETURNING uuid;")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// executes the query and gets the uuid returned
	var uuid uuid.UUID
	if err = stmt.QueryRow(nickname, registered).Scan(&uuid); err != nil {
		return nil, err
	}

	return &uuid, nil
}

// GetPlayer returns the player corresponding to the given uuid.
// It makes a SELECT query on the database, limited to one result.
func (b *Backend) GetPlayer(id *uuid.UUID) (*Player, error) {
	// executes a prepared SELECT request
	rows, err := b.db.Query("SELECT uuid, nickname, registered, faction FROM player WHERE uuid = $1 LIMIT 1;", id)
	if err != nil {
		return nil, err
	}

	// parses the result
	var player Player
	for rows.Next() {
		if err := rows.Scan(&player.UUID, &player.Nickname, &player.Registered, &player.Faction); err != nil {
			return nil, err
		}
	}

	return &player, nil
}

// DeletePlayer deletes a player corresponding to the given uuid from the database.
// It returns the uuid of the player deleted.
func (b *Backend) DeletePlayer(id *uuid.UUID) (*uuid.UUID, error) {
	// prepares the query
	stmt, err := b.db.Prepare("DELETE FROM player WHERE uuid = $1 returning uuid;")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// executes the query and gets the uuid returned
	var returnedUUID uuid.UUID
	if err = stmt.QueryRow(id).Scan(&returnedUUID); err != nil {
		return nil, err
	}

	return &returnedUUID, nil
}

// RegisterPlayer edits the "register" status of player corresponding to the given id.
// It returns the uuid of the player deleted.
func (b *Backend) RegisterPlayer(id *uuid.UUID, register bool) (*uuid.UUID, error) {
	stmt, err := b.db.Prepare("UPDATE player SET registered = $1 WHERE uuid = $2 returning uuid;")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var returnedUUID uuid.UUID
	if err = stmt.QueryRow(register, id).Scan(&returnedUUID); err != nil {
		return nil, err
	}

	return &returnedUUID, nil
}

// SetPlayerFaction sets faction to a player.
func (b *Backend) SetPlayerFaction(id, faction *uuid.UUID) (*uuid.UUID, error) {
	stmt, err := b.db.Prepare("UPDATE player SET faction = $1 WHERE uuid = $2 returning uuid;")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var returnedUUID uuid.UUID
	if err = stmt.QueryRow(faction, id).Scan(&returnedUUID); err != nil {
		return nil, err
	}

	return &returnedUUID, nil
}
