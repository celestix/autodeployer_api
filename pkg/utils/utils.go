package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateHex() string {
	var b = make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
