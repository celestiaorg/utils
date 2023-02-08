package misc

import (
	"crypto/rand"
)

const Charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"

// RandString generates a random string of any length
// This function relies on crypto/rand and thus is secure.
func RandString(n int) (string, error) {
	bz := make([]byte, n)
	_, err := rand.Read(bz)
	if err != nil {
		return "", err
	}

	const l = byte(len(Charset))
	for i, b := range bz {
		bz[i] = Charset[b%l]
	}

	return string(bz), nil
}
