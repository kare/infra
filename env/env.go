package env

import (
	"fmt"
	"strings"
)

// Env represents runtime environment.
type Env string

func (e Env) String() string {
	return string(e)
}

const (
	development = "development"
	staging     = "staging"
	production  = "production"
)

const (
	// Development is developers local environment not necessarily connected to Internet.
	Development = Env(development)
	// Staging describes production like environment such as testing or CI server.
	Staging = Env(staging)
	// Production describes the one and only production environment.
	Production = Env(production)
)

// EnvDefault is the default runtime environment.
var EnvDefault = Production

// MustParseEnv calls [ParseEnv] and on error returns [EnvDefault].
func MustParseEnv(s string) Env {
	if e, err := ParseEnv(s); err != nil {
		return EnvDefault
	} else {
		return e
	}
}

// ParseEnv parses environment from string and on unknown environment fails with error.
func ParseEnv(s string) (Env, error) {
	e := strings.ToLower(strings.TrimSpace(s))
	switch e {
	case development:
		return Development, nil
	case staging:
		return Staging, nil
	case production:
		return Production, nil
	default:
		return Env(""), fmt.Errorf("infra: unknown environment: '%s'", e)
	}
}
