package webhook

import (
	"net/http"
)

// Impelement the http.HandleFunc("/healthz", healthz)
//http.HandleFunc("/readyz", readyz)

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("okz"))
}

func Readyz(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("okz"))
}
