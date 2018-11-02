package test_test

import (
	"testing"

	"github.com/zeeraw/protogen/test"
)

func TestAssertEqual(t *testing.T) {
	test.AssertEqual(t, 1, 1)
	test.AssertEqual(t, "", "")
	test.AssertEqual(t, nil, nil)
}
