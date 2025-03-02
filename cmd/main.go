package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/vishalanarase/pod-webhook-mutator/pkg/config"
	"github.com/vishalanarase/pod-webhook-mutator/pkg/webhook"
)

// GetRouter returns a new router with the handlers for the healthz and readyz endpoints.
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", webhook.Healthz)
	r.HandleFunc("/readyz", webhook.Readyz)

	return r
}

func main() {
	// Load the configuration
	config := config.LoadConfig()
	fmt.Printf("Server address: %s\n", config.ServerAddress)

	// Server configuration
	server := &http.Server{
		Addr:      config.ServerAddress,
		TLSConfig: config.TLSConfig,
		Handler:   GetRouter(),
	}

	// Start the server
	fmt.Println("Starting server on :443")
	err := server.ListenAndServeTLS("certs/server.crt", "certs/server.key")
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println("Server started on :443")
}
