package server

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Route ... contains data to declare a route
type Route struct {
	Pattern     string
	Method      string
	Name        string
	HandlerFunc http.HandlerFunc
}

// Routes ... type to hold multiple routes
type Routes []Route

//HTTPResponder serves method to respond to http calls
type HTTPResponder interface {
	JSON(w http.ResponseWriter, code int, payload interface{})
	ERROR(w http.ResponseWriter, code int, err error)
}

// JSON responds to the request with the given code and payload
func (app *app) JSON(w http.ResponseWriter, code int, payload interface{}) {

	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)

}

// JSON responds to the request with the given code and payload
func (app *app) ERROR(w http.ResponseWriter, code int, err error) {
	response := []byte(err.Error())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}

// Logger function for http calls
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
