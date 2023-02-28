package infra

import (
	"fmt"
	"os"

	"kkn.fi/infra/env"
)

// Must checks for given err. If err is not nil panic is called, otherwise
// returns value T.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("infra: Must panic: %v", err))
	}
	return value
}

func getEnv() string {
	return os.Getenv("ENV")
}

// IsDevelopment checks environment variable ENV and returns true only if ENV value
// is `development`.
func IsDevelopment() bool {
	return getEnv() == env.Development.String()
}

// IsProduction checks the value of environment variable ENV and returns true
// if value is `production`.
func IsProduction() bool {
	return getEnv() == env.Production.String()
}

// IsCI returns true if environment variable CI has value true and false otherwise.
func IsCI() bool {
	return os.Getenv("CI") == "true"
}
