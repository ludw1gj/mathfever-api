package router

import (
	"github.com/gorilla/mux"
	"gitlab.com/ludw1gj/mathfever-api/src/handler"
	"net/http"
)

// cors Middleware
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// Load returns a router instance with routes
func Load() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.Use(cors)
	r.HandleFunc("/categories", handler.CategoriesHandler).Methods(http.MethodOptions, http.MethodGet)
	r.HandleFunc("/calculation/{slug}", handler.CalculationHandler).Methods(http.MethodOptions, http.MethodPost)
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)
	return r
}
