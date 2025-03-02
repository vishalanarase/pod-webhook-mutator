package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/vishalanarase/pod-webhook-mutator/pkg/config"
	"github.com/vishalanarase/pod-webhook-mutator/pkg/webhook"
)

func main() {
	// Load the configuration
	logrus.Info("Loading configuration")
	config, err := config.LoadConfig() // Assuming LoadConfig returns (config, error)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load configuration")
	}

	logrus.Info("Configuration loaded successfully")

	logrus.Infof("Server address: %s", config.ServerAddress)

	// Server configuration
	server := &http.Server{
		Addr:      config.ServerAddress,
		TLSConfig: config.TLSConfig,
		Handler:   webhook.GetRouter(),
	}

	// Start the server
	logrus.Info("Starting server on :443")
	err = server.ListenAndServeTLS("certs/server.crt", "certs/server.key")
	if err != nil {
		logrus.WithError(err).Error("Failed to start server")
		logrus.Fatal(err)
	}

	logrus.Info("Server started on :443")
}
