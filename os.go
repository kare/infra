package infra

import "os"

// GetenvDefault returns the value of environment variable called by key. If
// the value is undefined or in other words empty, then value of given
// defaultValue is returned.
func GetenvDefault(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}
