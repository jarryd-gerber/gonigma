package main

import (
	"fmt"
	"strings"

	"github.com/jarryd-gerber/gonigma/parts"
)

func main() {
	rotorI, _ := parts.CreateRotor("I", 0, 0)
	rotorII, _ := parts.CreateRotor("II", 0, 0)
	rotorIII, _ := parts.CreateRotor("III", 0, 0)
	reflector, _ := parts.CreateRotor("UKW-A", 0, 0)

	inputText := "E"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, char := range inputText {

		notch := rotorIII.Rotate()
		if notch {
			notch = rotorII.Rotate()
			if notch {
				rotorI.Rotate()
			}
		}

		input := strings.Index(alphabet, string(char))
		fmt.Printf("input is: %s\n", string(char))

		inputIII := rotorIII.GetRelativeInputPosition(input)
		pos, val := rotorIII.GetOuput(inputIII)
		fmt.Printf("output is: %s and %d\n", val, pos)

		inputII := rotorII.GetRelativeInputPosition(pos)
		pos, val = rotorII.GetOuput(inputII)
		fmt.Printf("output is: %s and %d\n", val, pos)

		inputI := rotorI.GetRelativeInputPosition(pos)
		pos, val = rotorI.GetOuput(inputI)
		fmt.Printf("output is: %s and %d\n", val, pos)

		pos, val = reflector.GetOuput(pos)
		fmt.Printf("output is: %s and %d\n", val, pos)

		inputI = rotorI.GetRelativeInputPosition(pos)
		pos, val = rotorI.GetReverseOutput(inputI)
		fmt.Printf("output is: %s and %d\n", val, pos)

		inputII = rotorII.GetRelativeInputPosition(pos)
		pos, val = rotorII.GetReverseOutput(inputII)
		fmt.Printf("output is: %s and %d\n", val, pos)

		inputIII = rotorIII.GetRelativeInputPosition(pos)
		_, val = rotorIII.GetReverseOutput(inputIII)
		fmt.Printf("output is: %s and %d\n", val, pos)
	}
}
