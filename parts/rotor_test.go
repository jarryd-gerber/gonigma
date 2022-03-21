package parts_test

import (
	"testing"

	"github.com/jarryd-gerber/gonigma/parts"
)

func TestRotorI(t *testing.T) {

	rotor := parts.CreateRotor("I", 0, 0)

	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "A", expect: "E"},
		{input: "B", expect: "K"},
		{input: "C", expect: "M"},
	}

	for _, scenario := range scenarios {
		if got := rotor.GetOutputValue(scenario.input); got != scenario.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected %q, got %q",
				scenario.input, got, scenario.expect)
		}
	}
}
