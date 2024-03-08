package utils

import (
	"crypto/rand"
	"math/big"
)

const codeLength = 10

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateCode() string {
	code := make([]rune, codeLength)
	max := big.NewInt(int64(len(alphabet)))

	for i := range code {
		idx, _ := rand.Int(rand.Reader, max)
		code[i] = alphabet[idx.Int64()]
	}

	return string(code)
}
