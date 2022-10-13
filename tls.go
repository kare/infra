package infra

import (
	"crypto/tls"

	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

func GetCertificateFuncFromFiles(certFile, keyFile string) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, err
		}
		return &cert, nil
	}
}

func GetCertificateFuncFromLetsEncrypt(m *autocert.Manager) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	return m.GetCertificate
}

func GetCertificateFunc(m *autocert.Manager, devCertFile, devKeyFile string) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	if IsProduction() {
		return GetCertificateFuncFromLetsEncrypt(m)
	}
	return GetCertificateFuncFromFiles(devCertFile, devKeyFile)
}

// TLSConfig returns secure TLS configuration for Internet server.
func TLSConfig(getCertificate func(*tls.ClientHelloInfo) (*tls.Certificate, error)) *tls.Config {
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
