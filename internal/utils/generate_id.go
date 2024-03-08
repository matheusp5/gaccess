package utils

import (
	"crypto/rand"
	"encoding/binary"
)

func GenerateId() uint32 {
	var id uint32
	binary.Read(rand.Reader, binary.BigEndian, &id)
	return id
}
