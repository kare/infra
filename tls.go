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
//
// Deprecated: GetCertificateFuncFromFiles is deprecated. Use GetCertificateFromFiles instead.
func GetCertificateFuncFromFiles(certFile, keyFile string) GetCertificate {
	return GetCertificateFromFiles(certFile, keyFile)
}

// GetCertificateFuncFromLetsEncrypt loads GetCertificate func from given Let's Encrypt Manager.
//
// Deprecated: GetCertificateFuncFromFiles is deprecated. Use GetCertificateFromLetsEncrypt instead.
func GetCertificateFuncFromLetsEncrypt(m *autocert.Manager) GetCertificate {
	return m.GetCertificate
}

// GetCertificateFunc loads certificates for development or production depending on return value
// of function [kkn.fi/infra#IsProduction].
//
// Deprecated: GetCertificateFunc is deprecated.
func GetCertificateFunc(m *autocert.Manager, devCertFile, devKeyFile string) GetCertificate {
	if IsProduction() {
		return GetCertificateFromLetsEncrypt(m)
	}
	return GetCertificateFuncFromFiles(devCertFile, devKeyFile)
}

// GoodTLSConfig configures a copy of given [tls.Config] with good settings.
//
// Deprecated: GooldTLSConfig is unused.
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

// GetCertificateFromFiles loads GetCertificate func from given TLS cert files.
func GetCertificateFromFiles(certFile, keyFile string) GetCertificate {
	return func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, err
		}
		return &cert, nil
	}
}

// GetCertificateFromLetsEncrypt loads GetCertificate func from given Let's Encrypt Manager.
func GetCertificateFromLetsEncrypt(m *autocert.Manager) GetCertificate {
	return m.GetCertificate
}

// TLSConfig returns secure TLS configuration for Internet server.
func TLSConfig(getCertificate GetCertificate) *tls.Config {
	nextProtos := []string{
		"h2",
		"http/1.1",
	}
	if IsProduction() {
		nextProtos = append(nextProtos, acme.ALPNProto)
	}
	conf := &tls.Config{
		GetCertificate: getCertificate,
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
		NextProtos: nextProtos,
	}
	return conf
}
