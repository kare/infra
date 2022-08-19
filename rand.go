package infra

import (
	"crypto/rand"
	"encoding/binary"
)

func RandInt64() (int64, error) {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	u := binary.LittleEndian.Uint64(b[:])
	r := int64(u)
	return r, nil
}
