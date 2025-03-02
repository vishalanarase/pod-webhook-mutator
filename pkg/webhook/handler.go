package webhook

import "github.com/gorilla/mux"

// GetRouter returns a new router with the handlers for the healthz and readyz endpoints.
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", Healthz)
	r.HandleFunc("/readyz", Readyz)

	return r
}
