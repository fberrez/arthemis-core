package backend

import "github.com/google/uuid"

// Faction is the struct containing informations
// about a faction.
type Faction struct {
	// UUID is the faction uuid.
	UUID uuid.UUID `json:"uuid" description:"Faction UUID."`

	// Name is the faction name.
	Name string `json:"name" description:"Faction name."`

	// Color is the faction color.
	Color int `json:"color" description:"Faction color."`
}

// AddFaction adds a new faction to the database.
// It returns the uuid of the faction added.
func (b *Backend) AddFaction(name string, color int) (*uuid.UUID, error) {
	// prepares the query
	stmt, err := b.db.Prepare("INSERT INTO faction(name, color) VALUES ($1, $2) RETURNING uuid;")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// executes the query and gets the uuid returned
	var uuid uuid.UUID
	if err = stmt.QueryRow(name, color).Scan(&uuid); err != nil {
		return nil, err
	}

	return &uuid, nil
}

// GetFaction returns the faction corresponding to the given uuid.
// It makes a SELECT query on the database, limited to one result.
func (b *Backend) GetFaction(id *uuid.UUID) (*Faction, error) {
	// executes a prepared SELECT request
	rows, err := b.db.Query("SELECT uuid, name, color FROM faction WHERE uuid = $1 LIMIT 1;", id)
	if err != nil {
		return nil, err
	}

	// parses the result
	var faction Faction
	for rows.Next() {
		if err := rows.Scan(&faction.UUID, &faction.Name, &faction.Color); err != nil {
			return nil, err
		}
	}

	return &faction, nil
}

// DeleteFaction deletes a faction corresponding to the given uuid from the database.
// It returns the uuid of the faction deleted.
func (b *Backend) DeleteFaction(id *uuid.UUID) (*uuid.UUID, error) {
	// prepares the query
	stmt, err := b.db.Prepare("DELETE FROM faction WHERE uuid = $1 returning uuid;")
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

// UpdateFaction updates the faction corresponding to the given id
// with the updated value of name and color.
// It returns the uuid of the updated faction.
func (b *Backend) UpdateFaction(id *uuid.UUID, name string, color int) (*uuid.UUID, error) {
	// prepares the query
	stmt, err := b.db.Prepare("UPDATE faction SET name = $1, color = $2 WHERE uuid = $3 returning uuid;")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// executes the query and gets the uuid returned
	var returnedUUID uuid.UUID
	if err = stmt.QueryRow(name, color, id).Scan(&returnedUUID); err != nil {
		return nil, err
	}

	return &returnedUUID, nil
}
