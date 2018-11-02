package test

import (
	"testing"
)

// AssertEqual asserts that the expected item is the same as the actual one
func AssertEqual(tb testing.TB, expected, actual interface{}) {
	tb.Helper()
	if expected != actual {
		tb.Fatalf("\n\texp: %#[1]v (%[1]T)\n\tgot: %#[2]v (%[2]T)\n", expected, actual)
	}
}
