package main

import (
	"fmt"

	"github.com/jarryd-gerber/gonigma/parts"
)

func main() {
	rotorI, _ := parts.CreateRotor("I", 0, 0)
	rotorII, _ := parts.CreateRotor("II", 0, 0)
	rotorIII, _ := parts.CreateRotor("III", 0, 0)
	reflector, _ := parts.CreateRotor("UKW-A", 0, 0)

	for i := 0; i < 25; i++ {
		notch := rotorI.Rotate()
		if notch {
			notch = rotorII.Rotate()
			if notch {
				rotorIII.Rotate()
			}
		}

		output := rotorI.GetInputPosition(
			rotorII.GetInputPosition(
				rotorIII.GetInputPosition(
					reflector.GetOuputPosition(
						rotorIII.GetOuputPosition(
							rotorII.GetOuputPosition(
								rotorI.GetOuputPosition(i)))))))

		fmt.Print(output)
	}
}
