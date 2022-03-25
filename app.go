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

	inputText := "AAAAA"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, char := range inputText {

		notch := rotorIII.Rotate()
		if notch {
			notch = rotorII.Rotate()
			if notch {
				rotorI.Rotate()
			}
		}

		value := strings.Index(alphabet, string(char))

		inputIII := rotorIII.GetRelativeInputPosition(value)
		outputIII := rotorIII.GetOuputPosition(inputIII)

		inputII := rotorII.GetRelativeInputPosition(outputIII)
		outputII := rotorII.GetOuputPosition(inputII)

		inputI := rotorI.GetRelativeInputPosition(outputII)
		outputI := rotorI.GetOuputPosition(inputI)

		outputR := reflector.GetOuputPosition(outputI)

		inputI = rotorI.GetRelativeInputPosition(outputR)
		outputI = rotorI.GetOuputPosition(inputI)

		inputII = rotorII.GetRelativeInputPosition(outputI)
		outputII = rotorII.GetOuputPosition(inputII)

		inputIII = rotorI.GetRelativeInputPosition(outputII)
		outputIII = rotorI.GetOuputPosition(inputIII)

		fmt.Print(string(alphabet[outputIII]))
	}
}
