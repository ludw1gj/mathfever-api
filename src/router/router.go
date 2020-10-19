package router

import (
	"github.com/gorilla/mux"
	"gitlab.com/ludw1gj/mathfever-api/src/handler"
	"net/http"
)

// Load returns a router instance with routes
func Load() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/categories", handler.CategoriesHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/calculation/{slug}", handler.CalculationHandler).Methods("POST", "OPTIONS")
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)
	r.Use(mux.CORSMethodMiddleware(r))
	return r
}
