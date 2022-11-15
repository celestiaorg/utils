package test

import (
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestRetry probes the Retry function
func TestRetry(t *testing.T) {
	// Define Functions
	returnError := errors.New("returned error")
	errorFunc := func() error { return returnError }
	successFunc := func() error { return nil }
	intermittentFunc := func() error {
		// Fail 25% of the time
		if rand.Intn(4) == 2 {
			return returnError
		}
		return nil
	}

	// Test Case: function always returns an error resulting in the full
	// execution of tries.
	//
	// Set a retry timer for 500 Milliseconds, this should finish before the
	// retry loop because the function always errors
	retryTimer := time.NewTimer(500 * time.Millisecond)

	// Set the Retry for 1 second total duration. We use milliseconds to
	// ensure the func is called multiple times
	err := Retry(10, 100*time.Millisecond, errorFunc)

	// Verify the Retry took longer than the timer
	select {
	case <-retryTimer.C:
	default:
		t.Error("Retry executed faster than expected")
	}

	// Ensure we have the error returned from the errorFunc
	assert.ErrorIs(t, err, returnError)

	// Clean up, stop timer
	retryTimer.Stop()

	// Test Case: function returns nil immediately
	//
	// Set a retry timer for 1 second, this should not finish before the
	// retry loop because the function returns nil
	retryTimer = time.NewTimer(time.Second)

	// Set the Retry for 10 second total duration, but we expect only one
	// iteration to occur since the function returns nil
	err = Retry(100, 100*time.Millisecond, successFunc)

	// Verify the Retry finished before the timer
	select {
	case <-retryTimer.C:
		t.Error("Retry executed slower than expected")
	default:
	}

	// Ensure we have a nil error returned from the successFunc
	assert.Nil(t, err)

	// Clean up, stop timer
	retryTimer.Stop()

	// Test Case: function sometimes returns an error, but is ultimately
	// expected to succeed
	//
	// Set a retry timer for 10 seconds, which is the same as the retry. The
	// Retry should finish first
	retryTimer = time.NewTimer(10 * time.Second)

	// Set the Retry for 10 second total duration, we give it 100 tries and
	// expect some to fail but ultimately it should return a nil
	err = Retry(100, 100*time.Millisecond, intermittentFunc)

	// Verify the Retry finished before the timer
	select {
	case <-retryTimer.C:
		t.Error("Retry executed slower than expected")
	default:
	}

	// Ensure we have a nil error returned from the successFunc
	assert.Nil(t, err)

	// Clean up, stop timer
	retryTimer.Stop()
}
