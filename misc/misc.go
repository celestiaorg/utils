package misc

import (
	"crypto/rand"
	"math/big"
)

const seed = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

// RandString generates a random string of any length
// and returns an error if an issue
// was encountered while doing so.
// This function relies on crypto/rand and thus is secure.
func RandString(n int) (string, error) {
	bz := make([]byte, n)
	_, err := rand.Read(bz)
	if err != nil {
		return "", err
	}

	const l = byte(len(seed))
	for i, b := range bz {
		bz[i] = seed[b%l]
	}

	return string(bz), nil
}
