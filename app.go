package main

import "github.com/jarryd-gerber/gonigma/parts"

func main() {
	rotorI, _ := parts.CreateRotor("I", 0, 0)
	rotorII, _ := parts.CreateRotor("II", 0, 0)
	rotorIII, _ := parts.CreateRotor("III", 0, 0)

	notch := rotorI.Rotate()
	if notch {
		notch = rotorII.Rotate()
		if notch {
			rotorIII.Rotate()
		}
	}
}
