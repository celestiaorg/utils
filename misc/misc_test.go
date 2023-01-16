package misc

import (
	"fmt"
	"regexp"
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

		re := regexp.MustCompile(fmt.Sprintf("^[%s]+$", Charset))

		assert.True(
			t,
			re.MatchString(randStr1),
			"the random string should not contain any characters outside the allowed charset",
		)

		assert.True(
			t,
			re.MatchString(randStr2),
			"the random string should not contain any characters outside the allowed charset",
		)
	}
}
