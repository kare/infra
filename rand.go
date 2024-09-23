package infra

import (
	"crypto/rand"
	"encoding/binary"
)

type randomReader interface {
	Read([]byte) (int, error)
}

type defaultRandomReader struct{}

func (d defaultRandomReader) Read(b []byte) (int, error) {
	return rand.Read(b)
}

var randReader randomReader = &defaultRandomReader{}

// RandInt64 returns cryptographically secure random number.
func RandInt64() (int64, error) {
	var b [8]byte
	if _, err := randReader.Read(b[:]); err != nil {
		return 0, err
	}
	u := binary.LittleEndian.Uint64(b[:])
	r := int64(u)
	return r, nil
}
