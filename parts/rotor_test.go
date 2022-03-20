package parts_test

import (
	"testing"

	"github.com/jarryd-gerber/gonigma/parts"
)

func TestRotor(t *testing.T) {

	rotor := parts.CreateRotor("I", 0)
	expected := "A"

	if got := rotor.GetOutput("A"); got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}
