package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler is responsible for defining a HTTP request route and corresponding handler.
type Handler struct {
	// Receives a route to modify, like adding path, methods, etc.
	Route func(r *mux.Route)

	// HTTP Handler
	Func http.HandlerFunc
}

// AddRoute adds the handler's route the to the router.
func (h Handler) AddRoute(r *mux.Router) {
	h.Route(r.NewRoute().HandlerFunc(h.Func))
}
func Greeter(prefix string) Handler {
	return Handler{
		Route: func(r *mux.Route) {
			r.Path("/greet/{name}").Methods("GET")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			name, ok := mux.Vars(r)["name"]
			if !ok || name == "" {
				name = "Champ"
			}
			_, err := w.Write([]byte(prefix + " " + name + "!"))
			if err != nil {
				log.Printf("Failed to write to response: %s\n", err)
			}
		},
	}
}
func main() {
	r := mux.NewRouter()
	complexDependency := "Hello" // ;)
	Greeter(complexDependency).AddRoute(r)
	log.Fatal(http.ListenAndServe(":1230", r))
}

