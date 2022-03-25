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

	inputText := "CTSYZO"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, char := range inputText {

		// notch := rotorI.Rotate()
		// if notch {
		// 	notch = rotorII.Rotate()
		// 	if notch {
		// 		rotorIII.Rotate()
		// 	}
		// }

		input := strings.Index(alphabet, string(char))
		//fmt.Printf("input is: %s\n", string(char))

		inputI := rotorI.GetRelativeInputPosition(input)
		pos, _ := rotorI.GetOuput(inputI)
		//fmt.Printf("rotorI: %s and %d\n", val, pos)

		inputII := rotorII.GetRelativeInputPosition(pos)
		pos, _ = rotorII.GetOuput(inputII)
		//fmt.Printf("rotorII: %s and %d\n", val, pos)

		inputIII := rotorIII.GetRelativeInputPosition(pos)
		pos, _ = rotorIII.GetOuput(inputIII)
		//fmt.Printf("rotorIII: %s and %d\n", val, pos)

		pos, _ = reflector.GetOuput(pos)
		//fmt.Printf("reflectorA: %s and %d\n", val, pos)

		inputIII = rotorIII.GetRelativeInputPosition(pos)
		pos, _ = rotorIII.GetReverseOutput(inputIII)
		//fmt.Printf("rotorIII: %s and %d\n", val, pos)

		inputII = rotorII.GetRelativeInputPosition(pos)
		pos, _ = rotorII.GetReverseOutput(inputII)
		//fmt.Printf("rotorII: %s and %d\n", val, pos)

		inputI = rotorI.GetRelativeInputPosition(pos)
		pos, _ = rotorI.GetReverseOutput(inputI)
		//fmt.Printf("rotorI: %s and %d\n", val, pos)

		fmt.Print(string(alphabet[pos]))
	}
}
