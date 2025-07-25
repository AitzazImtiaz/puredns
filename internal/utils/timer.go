package utils

import (
	"time"
)

// WithTimeout runs a function with a timeout.
func WithTimeout(f func() error, timeout time.Duration) error {
	done := make(chan error, 1)
	go func() {
		done <- f()
	}()

	select {
	case <-time.After(timeout):
		return &TimeoutError{}
	case err := <-done:
		return err
	}
}

type TimeoutError struct{}

func (e *TimeoutError) Error() string {
	return "operation timed out"
}
