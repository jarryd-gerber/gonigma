// Rotor #	ABCDEFGHIJKLMNOPQRSTUVWXYZ	Date Introduced	Model Name & Number
// I	EKMFLGDQVZNTOWYHXUSPAIBRCJ	1930	Enigma I
// II	AJDKSIRUXBLHWTMCQGZNPYFVOE	1930	Enigma I
// III	BDFHJLCPRTXVZNYEIWGAKMUSQO	1930	Enigma I
// IV	ESOVPZJAYQUIRHXLNFTGKDCMWB	December 1938	M3 Army
// V	VZBRGITYUPSDNHLXAWMJQOFECK	December 1938	M3 Army
// VI	JPGVOUMFYQBENHZRDKASXLICTW	1939	M3 & M4 Naval (FEB 1942)
// VII	NZJHGRCXMYSWBOUFAIVLPEKQDT	1939	M3 & M4 Naval (FEB 1942)
// VIII	FKQHTLXOCBJSPDZRAMEWNIUYGV	1939	M3 & M4 Naval (FEB 1942)

// Rotor	Notch	Effect
// I	Q	If rotor steps from Q to R, the next rotor is advanced
// II	E	If rotor steps from E to F, the next rotor is advanced
// III	V	If rotor steps from V to W, the next rotor is advanced
// IV	J	If rotor steps from J to K, the next rotor is advanced
// V	Z	If rotor steps from Z to A, the next rotor is advanced
// VI, VII, VIII	Z+M	If rotor steps from Z to A, or from M to N the next rotor is advanced
package parts

import (
	"fmt"
	"strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const minPosition = 0
const maxPosition = 25

var routes = map[string]string{
	"I":     "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	"II":    "AJDKSIRUXBLHWTMCQGZNPYFVOE",
	"III":   "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	"IV":    "ESOVPZJAYQUIRHXLNFTGKDCMWB",
	"V":     "VZBRGITYUPSDNHLXAWMJQOFECK",
	"UKW-A": "EJMZALYXVBWFCRQUONTSPIKHGD",
}

var notches = map[string]string{
	"I":   "Q",
	"II":  "E",
	"III": "V",
	"IV":  "J",
	"V":   "Z",
}

type Rotor struct {
	position      int
	ringSetting   int
	notchPosition string
	route         string
}

func CreateRotor(number string, startPosition int, ringSetting int) (*Rotor, error) {
	// Create a rotor and return a reference to it.
	if ringSetting > maxPosition {
		return nil, fmt.Errorf("ringSetting cannot be higher than 25")
	}

	if startPosition > maxPosition {
		return nil, fmt.Errorf("startPosition cannot be higher than 25")
	}

	return &Rotor{
		position:      startPosition,
		ringSetting:   ringSetting,
		notchPosition: notches[number],
		route:         routes[number],
	}, nil
}

func (r *Rotor) Rotate() bool {
	// rotates the rotor and returns True if notch position has been reached.
	if r.position == maxPosition {
		r.position = minPosition
	} else {
		r.position++
	}

	return r.position == strings.Index(alphabet, r.notchPosition)
}

func (r *Rotor) GetPosition() int {
	//Get the current position of a rotor
	return r.position
}

func (r *Rotor) GetRelativeInputPosition(outputPosition int) int {
	// Input of rotor relative to another rotors output.
	relativeInput := outputPosition + r.position

	if relativeInput > maxPosition {
		relativeInput -= maxPosition + 1
	}

	return relativeInput
}

func (r *Rotor) GetOuputPosition(inputPosition int) int {
	// Output of rotor for any given input.
	inputPosition += r.ringSetting

	outputPosition := strings.Index(alphabet, string(r.route[inputPosition]))

	return outputPosition
}
