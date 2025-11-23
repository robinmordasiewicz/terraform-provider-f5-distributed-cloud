// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/crypto/pkcs12"
)

// AuthMethod represents the authentication method used.
type AuthMethod string

const (
	AuthMethodCertificate AuthMethod = "certificate"
	AuthMethodToken       AuthMethod = "token"
)

// AuthConfig holds authentication configuration.
type AuthConfig struct {
	// Certificate-based authentication
	P12File     string
	P12Password string
	CertFile    string
	KeyFile     string

	// Token-based authentication
	APIToken string
}

// Authenticator provides authentication for HTTP requests.
type Authenticator interface {
	// Configure configures the HTTP client with authentication.
	Configure(client *http.Client) error
	// AddAuth adds authentication to an HTTP request.
	AddAuth(req *http.Request) error
	// Method returns the authentication method.
	Method() AuthMethod
}

// CertificateAuthenticator implements certificate-based authentication.
type CertificateAuthenticator struct {
	tlsConfig *tls.Config
}

// NewCertificateAuthenticator creates a new certificate authenticator.
func NewCertificateAuthenticator(config *AuthConfig) (*CertificateAuthenticator, error) {
	var cert tls.Certificate
	var err error

	if config.P12File != "" {
		cert, err = loadP12Certificate(config.P12File, config.P12Password)
		if err != nil {
			return nil, NewAuthenticationError(string(AuthMethodCertificate),
				fmt.Sprintf("failed to load P12 certificate from %s", config.P12File), err)
		}
	} else if config.CertFile != "" && config.KeyFile != "" {
		cert, err = tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
		if err != nil {
			return nil, NewAuthenticationError(string(AuthMethodCertificate),
				fmt.Sprintf("failed to load certificate pair from %s and %s", config.CertFile, config.KeyFile), err)
		}
	} else {
		return nil, NewAuthenticationError(string(AuthMethodCertificate),
			"no certificate credentials provided", nil)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	return &CertificateAuthenticator{
		tlsConfig: tlsConfig,
	}, nil
}

// Configure configures the HTTP client with TLS certificate authentication.
func (a *CertificateAuthenticator) Configure(client *http.Client) error {
	transport, ok := client.Transport.(*http.Transport)
	if !ok {
		transport = &http.Transport{}
		client.Transport = transport
	}
	transport.TLSClientConfig = a.tlsConfig
	return nil
}

// AddAuth adds authentication to an HTTP request (no-op for certificate auth).
func (a *CertificateAuthenticator) AddAuth(req *http.Request) error {
	// Certificate auth is handled at the TLS layer, no request modification needed
	return nil
}

// Method returns the authentication method.
func (a *CertificateAuthenticator) Method() AuthMethod {
	return AuthMethodCertificate
}

// TokenAuthenticator implements token-based authentication.
type TokenAuthenticator struct {
	token string
}

// NewTokenAuthenticator creates a new token authenticator.
func NewTokenAuthenticator(config *AuthConfig) (*TokenAuthenticator, error) {
	if config.APIToken == "" {
		return nil, NewAuthenticationError(string(AuthMethodToken),
			"no API token provided", nil)
	}

	return &TokenAuthenticator{
		token: config.APIToken,
	}, nil
}

// Configure configures the HTTP client (no-op for token auth).
func (a *TokenAuthenticator) Configure(client *http.Client) error {
	return nil
}

// AddAuth adds the Authorization header to an HTTP request.
func (a *TokenAuthenticator) AddAuth(req *http.Request) error {
	req.Header.Set("Authorization", "APIToken "+a.token)
	return nil
}

// Method returns the authentication method.
func (a *TokenAuthenticator) Method() AuthMethod {
	return AuthMethodToken
}

// loadP12Certificate loads a certificate from a PKCS#12 file.
// Supports P12 files with certificate chains (common with F5 XC certificates).
func loadP12Certificate(p12File string, password string) (tls.Certificate, error) {
	p12Data, err := os.ReadFile(p12File)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to read P12 file: %w", err)
	}

	// Use ToPEM which handles P12 files with certificate chains
	pemBlocks, err := pkcs12.ToPEM(p12Data, password)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to decode P12 file: %w", err)
	}

	// Separate key and certificate PEM blocks
	var keyPEM, certPEM []byte
	for _, block := range pemBlocks {
		pemData := pem.EncodeToMemory(block)
		if block.Type == "PRIVATE KEY" || block.Type == "RSA PRIVATE KEY" || block.Type == "EC PRIVATE KEY" {
			keyPEM = append(keyPEM, pemData...)
		} else if block.Type == "CERTIFICATE" {
			certPEM = append(certPEM, pemData...)
		}
	}

	if len(keyPEM) == 0 {
		return tls.Certificate{}, fmt.Errorf("no private key found in P12 file")
	}
	if len(certPEM) == 0 {
		return tls.Certificate{}, fmt.Errorf("no certificate found in P12 file")
	}

	return tls.X509KeyPair(certPEM, keyPEM)
}

// NewAuthenticator creates the appropriate authenticator based on config.
func NewAuthenticator(config *AuthConfig) (Authenticator, error) {
	// Prefer token authentication if provided
	if config.APIToken != "" {
		return NewTokenAuthenticator(config)
	}

	// Otherwise, try certificate authentication
	if config.P12File != "" || (config.CertFile != "" && config.KeyFile != "") {
		return NewCertificateAuthenticator(config)
	}

	return nil, NewAuthenticationError("unknown",
		"no valid authentication credentials provided: specify either api_token or certificate credentials", nil)
}
