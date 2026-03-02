package otp

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashOTP(otp string) string {
	hash := sha256.Sum256([]byte(otp))
	return hex.EncodeToString(hash[:])
}