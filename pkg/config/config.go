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

	certPath := "certs/server.crt"
	keyPath := "certs/server.key"
	caCertPath := "certs/ca.crt"

	crt, err := os.ReadFile(certPath)
	if err != nil {
		fmt.Println(err)
	}
	key, err := os.ReadFile(keyPath)
	if err != nil {
		fmt.Println(err)
	}
	caCert, err := os.ReadFile(caCertPath)
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
