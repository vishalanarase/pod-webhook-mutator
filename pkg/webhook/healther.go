package webhook

import "net/http"

// Healthz is a simple health check endpoint that returns a 200 OK status.
func Healthz(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("okz"))
	return
}

// Readyz is a simple readiness check endpoint that returns a 200 OK status.
func Readyz(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("okz"))
	return
}
