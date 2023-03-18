package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
)

func ConvertBytesToHexString(bytes []byte) string {
	hexStrings := make([]string, len(bytes))
	for i, b := range bytes {
		hexStrings[i] = fmt.Sprintf("0x%02x", b)
	}
	return strings.Join(hexStrings, ", ")
}

func GenerateRandomBytes(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	return b, err
}