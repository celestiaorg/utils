package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	for i := 0; i < 100; i++ {
		randStr1, err1 := RandString(10)
		randStr2, err2 := RandString(10)

		assert.Nil(t, err1)
		assert.Nil(t, err2)

		assert.Equal(t, 10, len(randStr1))
		assert.Equal(t, 10, len(randStr2))

		assert.NotEqual(t, randStr1, randStr2)
	}
}
