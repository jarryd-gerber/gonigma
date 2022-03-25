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
		{number: "I", ringSetting: 0, startPosition: 4, expect: 11},
		{number: "I", ringSetting: 0, startPosition: 5, expect: 6},
		{number: "I", ringSetting: 0, startPosition: 6, expect: 3},
		{number: "I", ringSetting: 0, startPosition: 7, expect: 16},
		{number: "I", ringSetting: 0, startPosition: 8, expect: 21},
		{number: "I", ringSetting: 0, startPosition: 9, expect: 25},
		{number: "I", ringSetting: 0, startPosition: 10, expect: 13},
		{number: "I", ringSetting: 0, startPosition: 11, expect: 19},
		{number: "I", ringSetting: 0, startPosition: 12, expect: 14},
		{number: "I", ringSetting: 0, startPosition: 13, expect: 22},
		{number: "I", ringSetting: 0, startPosition: 14, expect: 24},
		{number: "I", ringSetting: 0, startPosition: 15, expect: 7},
		{number: "I", ringSetting: 0, startPosition: 16, expect: 23},
		{number: "I", ringSetting: 0, startPosition: 17, expect: 20},
		{number: "I", ringSetting: 0, startPosition: 18, expect: 18},
		{number: "I", ringSetting: 0, startPosition: 19, expect: 15},
		{number: "I", ringSetting: 0, startPosition: 20, expect: 0},
		{number: "I", ringSetting: 0, startPosition: 21, expect: 8},
		{number: "I", ringSetting: 0, startPosition: 22, expect: 1},
		{number: "I", ringSetting: 0, startPosition: 23, expect: 17},
		{number: "I", ringSetting: 0, startPosition: 24, expect: 2},
		{number: "I", ringSetting: 0, startPosition: 25, expect: 9},
	}

	for _, scenario := range scenarios {
		rotor, _ := parts.CreateRotor(
			scenario.number, scenario.startPosition, scenario.ringSetting)

		if got := rotor.GetOuputPosition(scenario.startPosition); got != scenario.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected %d, got %d",
				scenario.startPosition, scenario.expect, got)
		}
	}
}

func TestRotate(t *testing.T) {
	startPosition := 3
	endPosition := 4
	rotor, _ := parts.CreateRotor("III", startPosition, 0)

	rotor.Rotate()

	if got := rotor.GetPosition(); got != endPosition {
		t.Errorf("Did not get expected result. Expected %d, got %d",
			endPosition, endPosition)
	}
}

func TestGetRelativeInputPosition(t *testing.T) {

	scenarios := []struct {
		outputPosition int
		expect         int
	}{
		{outputPosition: 1, expect: 17},
		{outputPosition: 10, expect: 0},
	}

	rotor, _ := parts.CreateRotor("I", 16, 0)

	for _, scenario := range scenarios {
		if got := rotor.GetRelativeInputPosition(scenario.outputPosition); got != scenario.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected %d, got %d",
				scenario.outputPosition, scenario.expect, got)
		}
	}
}

func TestGetOutputPosition(t *testing.T) {

	scenarios := []struct {
		inputPosition int
		ringSetting   int
		expect        int
	}{
		{inputPosition: 1, ringSetting: 0, expect: 10},
		{inputPosition: 1, ringSetting: 1, expect: 12},
	}

	for _, scenario := range scenarios {
		rotor, _ := parts.CreateRotor("I", 0, scenario.ringSetting)

		if got := rotor.GetOuputPosition(scenario.inputPosition); got != scenario.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected %d, got %d",
				scenario.inputPosition, scenario.expect, got)
		}
	}
}
