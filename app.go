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

	inputText := "ENIGMA"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, char := range inputText {

		notch := rotorI.Rotate()
		if notch {
			notch = rotorII.Rotate()
			if notch {
				rotorIII.Rotate()
			}
		}

		value := strings.Index(alphabet, string(char))

		outputIII := rotorIII.GetOuputPosition(value)

		inputII := rotorII.GetRelativeInputPosition(outputIII)
		outputII := rotorII.GetOuputPosition(inputII)

		inputI := rotorI.GetRelativeInputPosition(outputII)
		outputI := rotorI.GetOuputPosition(inputI)

		outputR := reflector.GetOuputPosition(outputI)

		outputI = rotorI.GetRelativeOutputPosition(outputR)
		inputI = rotorI.GetInputPosition(outputI)

		outputII = rotorII.GetRelativeOutputPosition(inputI)
		inputII = rotorII.GetInputPosition(outputII)

		outputIII = rotorI.GetRelativeOutputPosition(inputII)
		inputIII := rotorI.GetInputPosition(outputIII)

		fmt.Print(string(alphabet[inputIII]))
	}
}
