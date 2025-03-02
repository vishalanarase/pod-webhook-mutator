package webhook

import "github.com/gorilla/mux"

// GetRouter returns a new router with the handlers for the healthz and readyz endpoints.
func GetRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", Healthz).Methods("GET")
	r.HandleFunc("/readyz", Readyz).Methods("GET")
	r.HandleFunc("/validate", Validate).Methods("POST")
	r.HandleFunc("/mutate", Mutate).Methods("POST")

	return r
}
