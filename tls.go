package infra

import "crypto/tls"

// TLSConfig returns TLS configuration for kkn.fi server.
func TLSConfig(getCert func() func(*tls.ClientHelloInfo) (*tls.Certificate, error)) *tls.Config {
	conf := &tls.Config{
		GetCertificate: getCert(),
		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
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
			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients
			//tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			//tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		},
	}
	return conf
}
