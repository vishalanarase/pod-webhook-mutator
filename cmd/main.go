package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/vishalanarase/pod-webhook-mutator/pkg/config"
	"github.com/vishalanarase/pod-webhook-mutator/pkg/webhook"
)

func main() {
	// Load the configuration
	config := config.LoadConfig()
	fmt.Printf("Server address: %s\n", config.ServerAddress)

	// Server configuration
	server := &http.Server{
		Addr:      config.ServerAddress,
		TLSConfig: config.TLSConfig,
		Handler:   webhook.GetRouter(),
	}

	// Start the server
	fmt.Println("Starting server on :443")
	err := server.ListenAndServeTLS("certs/server.crt", "certs/server.key")
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println("Server started on :443")
}
