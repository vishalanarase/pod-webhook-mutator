package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

type Config struct {
	ServerAddress string
	TLSConfig     *tls.Config
}

func LoadConfig() *Config {
	address := os.Getenv("WEBHOOK_SERVER_ADDRESS")
	if address == "" {
		address = ":443" // Default to port 443
	}

	return &Config{
		ServerAddress: address,
		TLSConfig:     GetTLSConfig(),
	}
}

// GetTLSConfig returns a TLS configuration for the server.
func GetTLSConfig() *tls.Config {
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

	return tlsConfig
}
