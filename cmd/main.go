package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/vishalanarase/pod-webhook-mutator/pkg/webhook"
)

func main() {
	fmt.Println("Hello, World!")

	// Load the Server certificate and key
	crt, err := os.ReadFile("certs/server.crt")
	if err != nil {
		fmt.Println(err)
	}
	key, err := os.ReadFile("certs/server.key")
	if err != nil {
		fmt.Println(err)
	}

	// Create a cert pool and add the webhook's CA cert to it
	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		fmt.Println(err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		ClientCAs: certPool,
		//ClientAuth: tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{
			{
				Certificate: [][]byte{crt},
				PrivateKey:  key,
			},
		},
	}

	// Add the handlers for the healthz and readyz endpoints
	http.HandleFunc("/healthz", webhook.Healthz)
	http.HandleFunc("/readyz", webhook.Readyz)

	// Server configuration
	server := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
	}

	fmt.Println("Starting server on :443")
	err = server.ListenAndServeTLS("certs/server.crt", "certs/server.key")
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println("Server started on :443")
}
