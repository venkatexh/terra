package token

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateToken() (raw string, hash string, err error) {
	bytes := make([]byte, 32)

	_, err = rand.Read(bytes)
	if err != nil {
		return "", "", err
	}

	raw = hex.EncodeToString(bytes)

	hashBytes := sha256.Sum256([]byte(raw))
	hash = hex.EncodeToString(hashBytes[:])

	return raw, hash, nil
}