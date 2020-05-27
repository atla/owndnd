package server

import (
	"encoding/json"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"net/http"
	"strings"

	"encoding/base64"

	"github.com/buger/jsonparser"

	"github.com/atla/owndnd/pkg/db"
	"github.com/atla/owndnd/pkg/server/handler"
	"github.com/atla/owndnd/pkg/service"

	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

// App ... main application structure
type App interface {
	Run()
}

type app struct {

	// generic app base
	Router *gin.Engine
	db     *db.Client
	// owndnd specific
	facade service.Facade
}

// NewApp returns an application instance
// this is the primary stateless server providing an API interface
func NewApp() App {

	db := db.New(os.Getenv("MONGODB_DATABASE"))

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return &app{
		db:     db,
		Router: r,
		facade: service.NewFacade(db),
	}
}

/// AUTH0 handling
type jwks struct {
	Keys []webKeys `json:"keys"`
}
type webKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(os.Getenv("AUTH0_WK_JWKS"))

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}

func (app *app) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		r := c.Request

		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				// Verify 'aud' claim
				aud := os.Getenv("AUTH0_AUDIENCE")
				checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
				if !checkAud {
					return token, errors.New("Invalid audience")
				}
				// Verify 'iss' claim
				iss := os.Getenv("AUTH0_DOMAIN")
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
				if !checkIss {
					return token, errors.New("Invalid issuer")
				}

				cert, err := getPemCert(token)
				if err != nil {
					panic(err.Error())
				}

				result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
				return result, nil
			})

		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)

			c.AbortWithStatus(401)

		} else {

			// set userid if not already in context
			if _, ok := c.Get("userid"); !ok {
				splitted := strings.Split(token.Raw, ".")
				if decoded, err := base64.RawStdEncoding.DecodeString(splitted[1]); err == nil {
					if sub, err := jsonparser.GetString(decoded, "sub"); err == nil {
						c.Set("userid", sub)
					} else {
						log.WithError(err).Error("Could not get sub part from JSON")
					}
				} else {
					//TODO: remove token logging
					log.WithError(err).WithField("RawToken", token.Raw).Error("Could not decode token part")
				}
			}

			if _, ok := c.Get("user"); !ok {
				if id, exists := c.Get("userid"); exists {
					if user, err := app.facade.UsersService().FindOrCreateNewUser(id.(string)); err == nil {
						log.WithField("UserID", user.ID).Debug("Set user in Context")
						c.Set("user", user)
					}
				}
			}

			c.Set("token", token)
			c.Next()
		}
	}
}

// SetupRoutes ... Configures the routes
func (app *app) setupRoutes() {

	r := app.Router

	csh := &handler.CharactersHandler{
		app.facade.CharacterSheetsService(),
	}

	usr := &handler.UsersHandler{
		app.facade.UsersService(),
	}

	rooms := &handler.RoomsHandler{
		app.facade.RoomsService(),
	}

	r.Static("/app", "public/app/public")

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "API is up and running")
	})

	apiv1 := r.Group("/api/")
	//apiv1.Use(app.authMiddleware())
	{
		apiv1.GET("characters", csh.GetCharacters)
		apiv1.POST("characters", csh.PostCharacter)
		apiv1.GET("characters/:id", csh.GetCharacterByID)
		apiv1.DELETE("characters/:id", csh.DeleteCharacterByID)
		apiv1.PUT("characters/:id", csh.UpdateCharacterByID)

		apiv1.GET("rooms", rooms.GetRooms)
		apiv1.POST("rooms", rooms.PostRoom)
		apiv1.PUT("rooms/:id", rooms.PutRoom)

		apiv1.GET("user", usr.GetUser)
		apiv1.PUT("user", usr.UpdateUser)
	}

	/*

			//charactersheets
			Route{"/api/", "GET", "Get all charactersheets", app.authMiddleware(handler(csh.GetCharacterSheets))},
			Route{"/api/charactersheets/{id}", "GET", "Get a charactersheet by id", handler(csh.GetCharacterSheetByID)},
			Route{"/api/charactersheets/{id}", "DELETE", "Delete charactersheet by id", app.authMiddleware(handler(csh.DeleteCharacterSheetByID))},
			Route{"/api/charactersheets/{id}", "PUT", "Update charactersheet by id", handler(csh.UpdateCharacterSheetByID)},
			Route{"/api/charactersheets", "POST", "Create a new charactersheet", app.authMiddleware(handler(csh.PostCharacterSheet))},

			// parties
			Route{"/api/parties", "POST", "Create a new party", handler(ph.CreateParty)},
			Route{"/api/parties", "GET", "Get parties", handler(ph.GetParties)},
			Route{"/api/parties/{id}", "GET", "Get party by ID", handler(ph.GetPartyByID)},
			Route{"/api/parties/{id}", "PUT", "Update a party", handler(ph.UpdateParty)},
			Route{"/api/parties/{id}", "DELETE", "Delete a party", handler(ph.DeleteParty)},


			Route{"/health", "GET", "health check", statusHandler},
		}
	*/
}

// Run ... starts the server
func (app *app) Run() {

	app.db.Connect(os.Getenv("MONGODB_CONNECTION_STRING"))
	app.setupRoutes()

	port := 8010
	server := fmt.Sprintf("0.0.0.0:%v", port)

	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(app.Router)

	log.WithField("PORT", port).Info(fmt.Sprintf("ownDnD Server is running, listening on port %v", port))
	log.Fatal(http.ListenAndServe(server, corsHandler))
}
