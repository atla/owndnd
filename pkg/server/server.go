package server

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"net/http"

	"github.com/atla/owndnd/pkg/db"
	"github.com/atla/owndnd/pkg/service"
	"github.com/gorilla/mux"
)

// App ... main application structure
type App interface {
	Run()
}

type app struct {

	// generic app base
	Router *mux.Router
	routes Routes
	db     *db.Client

	// owndnd specific
	facade service.Facade
}

// NewApp returns an application instance
// this is the primary stateless server providing an API interface
func NewApp() App {

	db := db.New(os.Getenv("MONGODB_DATABASE"))

	return &app{
		db:     db,
		facade: service.NewFacade(db),
	}
}

// SetupRoutes ... Configures the routes
func (app *app) setupRoutes() {

	var statusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is up and running"))
	})

	csh := NewCharacterSheetHandler(app.facade.CharacterSheetsService(), app)
	ph := NewPartiesHandler(app.facade.PartiesService(), app)

	app.routes = Routes{
		//charactersheets
		Route{"/api/charactersheets", "GET", "Get all charactersheets", csh.GetCharacterSheets},
		Route{"/api/charactersheets/{id}", "GET", "Get a charactersheet by id", csh.GetCharacterSheetByID},
		Route{"/api/charactersheets/{id}", "DELETE", "Delete charactersheet by id", csh.DeleteCharacterSheetByID},
		Route{"/api/charactersheets/{id}", "PUT", "Update charactersheet by id", csh.UpdateCharacterSheetByID},
		Route{"/api/charactersheets", "POST", "Create a new charactersheet", csh.PostCharacterSheet},

		// parties
		Route{"/api/parties", "POST", "Create a new party", ph.CreateParty},
		Route{"/api/parties", "GET", "Get parties", ph.GetParties},
		Route{"/api/parties/{id}", "GET", "Get party by ID", ph.GetPartyByID},
		Route{"/api/parties/{id}", "PUT", "Update a party", ph.UpdateParty},
		Route{"/api/parties/{id}", "DELETE", "Delete a party", ph.DeleteParty},

		Route{"/health", "GET", "health check", statusHandler},
	}

	// wrap all routes in logger
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range app.routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	app.Router = router

	// also setup static serving
	//app.Router.PathPrefix("/creator").Handler(http.FileServer(http.Dir("public/creator/dist")))
	//app.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("public/viewer/public")))
}

// Run ... starts the server
func (app *app) Run() {

	app.db.Connect(os.Getenv("MONGODB_CONNECTION_STRING"))
	app.setupRoutes()

	fmt.Println("ownDnD Server is running, listening on port 8010")
	log.Fatal(http.ListenAndServe("0.0.0.0:8010", app.Router))
}
