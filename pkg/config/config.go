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

func LoadConfig() (*Config, error) {
	address := os.Getenv("WEBHOOK_SERVER_ADDRESS")
	if address == "" {
		address = ":443" // Default to port 443
	}

	tlsConfig, err := GetTLSConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get TLS config: %w", err)
	}

	return &Config{
		ServerAddress: address,
		TLSConfig:     tlsConfig,
	}, nil
}

// GetTLSConfig returns a TLS configuration for the server.
func GetTLSConfig() (*tls.Config, error) {
	// Load the Server certificate and key

	certPath := "certs/server.crt"
	keyPath := "certs/server.key"
	caCertPath := "certs/ca.crt"

	crt, err := os.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read server certificate: %w", err)
	}
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read server key: %w", err)
	}
	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate: %w", err)
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

	return tlsConfig, nil
}
