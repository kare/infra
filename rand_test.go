package infra

import (
	"errors"
	"testing"
)

func TestRandHappy(t *testing.T) {
	_, err := RandInt64()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

type okRandReader struct{}

func (*okRandReader) Read(b []byte) (int, error) {
	return len(b), nil
}

func TestRand(t *testing.T) {
	randReader = &okRandReader{}
	r, err := RandInt64()
	if r != 0 {
		t.Errorf("expecting %d, but got %d", 0, r)
	}
	if err != nil {
		t.Errorf("expecting nil error, but got: '%v'", err)
	}
}

type errRandReader struct{}

var msg = "mocked i/o error while reading /dev/random"

func (*errRandReader) Read(b []byte) (int, error) {
	return 0, errors.New(msg)
}

func TestRandError(t *testing.T) {
	randReader = &errRandReader{}
	_, err := RandInt64()
	if err == nil {
		t.Errorf("expecting error, but got nil")
	}
	if err.Error() != msg {
		t.Errorf("expecting '%s', got '%s'", msg, err.Error())
	}
}
