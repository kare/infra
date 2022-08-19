package infra

import "fmt"

// Must checks for given err. If err is not nil panic is called, otherwise
// returns value T.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("infra: Must panic: %v", err))
	}
	return value
}
