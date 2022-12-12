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
	seedlen := int64(len(seed))
	randstr := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(seedlen))
		if err != nil {
			return "", err
		}
		randstr[i] = seed[num.Int64()]
	}

	return string(randstr), nil
}
