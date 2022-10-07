package rs

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	lengthExpected := 10

	if got := GenerateRandomString(lengthExpected); len(got) != lengthExpected {
		t.Errorf("GenerateRandomString(%d) expected %d but got %d", lengthExpected, lengthExpected, len(got))
	}
}
