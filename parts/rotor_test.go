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
		{number: "I", ringSetting: 4, startPosition: 0, expect: 11},
		{number: "I", ringSetting: 5, startPosition: 0, expect: 6},
		{number: "I", ringSetting: 6, startPosition: 0, expect: 3},
		{number: "I", ringSetting: 7, startPosition: 0, expect: 16},
		{number: "I", ringSetting: 8, startPosition: 0, expect: 21},
		{number: "I", ringSetting: 9, startPosition: 0, expect: 25},
		{number: "I", ringSetting: 10, startPosition: 0, expect: 13},
		{number: "I", ringSetting: 11, startPosition: 0, expect: 19},
		{number: "I", ringSetting: 12, startPosition: 0, expect: 14},
		{number: "I", ringSetting: 13, startPosition: 0, expect: 22},
		{number: "I", ringSetting: 14, startPosition: 0, expect: 24},
		{number: "I", ringSetting: 15, startPosition: 0, expect: 7},
		{number: "I", ringSetting: 16, startPosition: 0, expect: 23},
		{number: "I", ringSetting: 17, startPosition: 0, expect: 20},
		{number: "I", ringSetting: 18, startPosition: 0, expect: 18},
		{number: "I", ringSetting: 19, startPosition: 0, expect: 15},
		{number: "I", ringSetting: 20, startPosition: 0, expect: 0},
		{number: "I", ringSetting: 21, startPosition: 0, expect: 8},
		{number: "I", ringSetting: 22, startPosition: 0, expect: 1},
		{number: "I", ringSetting: 23, startPosition: 0, expect: 17},
		{number: "I", ringSetting: 24, startPosition: 0, expect: 2},
		{number: "I", ringSetting: 25, startPosition: 0, expect: 9},
	}

	for _, scenario := range scenarios {
		rotor, _ := parts.CreateRotor(
			scenario.number, scenario.startPosition, scenario.ringSetting)

		if got := rotor.GetOuputPosition(rotor.GetPosition()); got != scenario.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected %d, got %d",
				scenario.ringSetting, scenario.expect, got)
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
		rotor, _ := parts.CreateRotor(
			scenario.number, scenario.startPosition, scenario.ringSetting)

		if got := rotor.GetOuputPosition(rotor.GetPosition()); got != scenario.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected %d, got %d",
				scenario.number, scenario.expect, got)
		}
	}
}
