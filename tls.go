package infra

import (
	"crypto/tls"

	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

// GetCertificate returns a Certificate based on the given ClientHelloInfo.
type GetCertificate func(*tls.ClientHelloInfo) (*tls.Certificate, error)

// Deprecated: GetDevelopmentCert is deprecated. Use GetCertificateFuncFromFiles instead.
func GetDevelopmentCert(certFile, keyFile string) GetCertificate {
	return func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, err
		}
		return &cert, nil
	}
}

// GetCertificateFuncFromFiles loads GetCertificate func from given TLS cert files.
func GetCertificateFuncFromFiles(certFile, keyFile string) GetCertificate {
	return func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, err
		}
		return &cert, nil
	}
}

// GetCertificateFuncFromLetsEncrypt loads GetCertificate func from given Let's Encrypt Manager.
func GetCertificateFuncFromLetsEncrypt(m *autocert.Manager) GetCertificate {
	return m.GetCertificate
}

// GetCertificateFunc loads certificates for development or production depending on return value
// of function [kkn.fi/infra#IsProduction].
func GetCertificateFunc(m *autocert.Manager, devCertFile, devKeyFile string) GetCertificate {
	if IsProduction() {
		return GetCertificateFuncFromLetsEncrypt(m)
	}
	return GetCertificateFuncFromFiles(devCertFile, devKeyFile)
}

// GoodTLSConfig conofigures given [tls.Config] with good settings.
func GoodTLSConfig(src *tls.Config) *tls.Config {
	result := src
	// Only use curves which have assembly implementations.
	result.CurvePreferences = []tls.CurveID{
		tls.CurveP256,
		tls.X25519,
	}
	result.MinVersion = tls.VersionTLS12
	result.CipherSuites = []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
	}
	return result
}

// TLSConfig returns secure TLS configuration for Internet server.
func TLSConfig(getCertificate GetCertificate) *tls.Config {
	conf := &tls.Config{
		GetCertificate: getCertificate,
		// Causes servers to use Go's default ciphersuite preferences, which
		// are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations.
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		NextProtos: []string{
			"h2",
			"http/1.1",
			acme.ALPNProto,
		},
	}
	return conf
}
