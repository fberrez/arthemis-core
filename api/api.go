package api

import (
	"net/http"
	"time"

	"bitbucket.org/arthemismc/core/backend"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/loopfz/gadgeto/tonic/utils/jujerr"
	log "github.com/sirupsen/logrus"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

// API contains each part of the API settings.
type (
	API struct {
		// fizz is the http server.
		fizz *fizz.Fizz

		// backend is the application backend.
		backend *backend.Backend
	}
)

// New parses the config file and initializes the new API.
func New() (*API, error) {
	log.Info("Initializing API")

	f := fizz.New()
	b, err := backend.New()
	if err != nil {
		return nil, err
	}

	api := &API{
		fizz:    f,
		backend: b,
	}

	// API informations
	infos := &openapi.Info{
		Title:       "Arthemis Core - Manage your Minecraft server",
		Description: "Arthemis is an API which helps you to manage your Minecraft server",
		Version:     "0.0.1",
	}

	// Defines groups of routes
	playersGroup := f.Group("/players", "Player", "Group of routes to interact with players.")
	factionsGroup := f.Group("/factions", "Faction", "Group of routes to interact with factions.")
	unsecuredGroup := f.Group("/unsecured", "Unsecured", "Group of unsecured routes.")

	// Defines Unsecured group's routes
	unsecuredGroup.GET("/openapi.json", []fizz.OperationOption{
		fizz.Summary("Generates a Swagger documentation in JSON"),
		fizz.Description("Returns a Swagger JSON containing all informations about the API."),
	}, f.OpenAPI(infos, "json"))

	unsecuredGroup.GET("/health", []fizz.OperationOption{
		fizz.Summary("Returns the health status of the API."),
	}, tonic.Handler(api.health, http.StatusOK))

	// Defines player group's routes
	playersGroup.POST("/", []fizz.OperationOption{
		fizz.Summary("Adds a new player"),
		fizz.Description("Adds a new player to the database. It also checks if it does not exist already."),
	}, tonic.Handler(api.addPlayer, http.StatusOK))

	playersGroup.GET("/:id", []fizz.OperationOption{
		fizz.Summary("Gets informations about a player"),
	}, tonic.Handler(api.getPlayer, http.StatusOK))

	playersGroup.DELETE("/:id", []fizz.OperationOption{
		fizz.Summary("Deletes an existing player corresponding to the given id."),
	}, tonic.Handler(api.deletePlayer, http.StatusOK))

	playersGroup.PUT("/register/:id", []fizz.OperationOption{
		fizz.Summary("Registers (or unregisters) a player."),
	}, tonic.Handler(api.registerPlayer, http.StatusOK))

	playersGroup.PUT("/faction/:id", []fizz.OperationOption{
		fizz.Summary("Sets faction to a player."),
	}, tonic.Handler(api.setPlayerFaction, http.StatusOK))

	// Defines faction group's routes
	factionsGroup.POST("/", []fizz.OperationOption{
		fizz.Summary("Adds a new faction"),
		fizz.Description("Adds a new faction to the database. It also checks if it does not exist already."),
	}, tonic.Handler(api.addFaction, http.StatusOK))

	factionsGroup.GET("/:id", []fizz.OperationOption{
		fizz.Summary("Gets informations about a faction"),
	}, tonic.Handler(api.getFaction, http.StatusOK))

	factionsGroup.DELETE("/:id", []fizz.OperationOption{
		fizz.Summary("Deletes an existing faction corresponding to the given id."),
	}, tonic.Handler(api.deleteFaction, http.StatusOK))

	factionsGroup.PUT("/:id", []fizz.OperationOption{
		fizz.Summary("Updates informations about a faction."),
	}, tonic.Handler(api.updateFaction, http.StatusOK))

	// Sets the error hook
	tonic.SetErrorHook(jujerr.ErrHook)

	return api, nil
}

// ServeHTTP is the implementation of http.Handler.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.WithFields(log.Fields{
		"start_time":  start.Unix(),
		"remote_addr": r.RemoteAddr,
		"request":     r.RequestURI,
	}).Info("Request received.")

	a.fizz.ServeHTTP(w, r)
}
