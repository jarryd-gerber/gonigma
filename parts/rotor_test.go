package parts_test

import (
	"testing"

	"github.com/jarryd-gerber/gonigma/parts"
)

func TestRingSettings(t *testing.T) {

	scenarios := []struct {
		number        string
		ringSetting   int
		startPosition int
		expect        int
	}{
		{number: "I", ringSetting: 0, startPosition: 0, expect: 4},
		{number: "I", ringSetting: 1, startPosition: 0, expect: 10},
		{number: "I", ringSetting: 2, startPosition: 0, expect: 12},
		{number: "I", ringSetting: 3, startPosition: 0, expect: 5},
	}

	for _, scenario := range scenarios {
		rotor := parts.CreateRotor(
			scenario.number, scenario.startPosition, scenario.ringSetting)

		if got := rotor.GetOuputPosition(); got != scenario.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected %d, got %d",
				scenario.number, scenario.expect, got)
		}
	}
}

func TestStartPositions(t *testing.T) {

	scenarios := []struct {
		number        string
		ringSetting   int
		startPosition int
		expect        int
	}{
		{number: "I", ringSetting: 0, startPosition: 0, expect: 4},
		{number: "I", ringSetting: 0, startPosition: 1, expect: 10},
		{number: "I", ringSetting: 0, startPosition: 2, expect: 12},
		{number: "I", ringSetting: 0, startPosition: 3, expect: 5},
	}

	for _, scenario := range scenarios {
		rotor := parts.CreateRotor(
			scenario.number, scenario.startPosition, scenario.ringSetting)

		if got := rotor.GetOuputPosition(); got != scenario.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected %d, got %d",
				scenario.number, scenario.expect, got)
		}
	}
}
