package otp

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GenerateOTP() (string, error) {

	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(n.Int64())), nil
}
