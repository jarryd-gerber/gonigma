package parts

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const minPosition = 0
const maxPosition = 25

type Contact struct {
	input  string
	output string
}

type Rotor struct {
	position      int
	ringSetting   int
	notchPosition string
	Contacts      []Contact
}

func createContacts(number string) []Contact {
	// Rotor #	ABCDEFGHIJKLMNOPQRSTUVWXYZ	Date Introduced	Model Name & Number
	// I	EKMFLGDQVZNTOWYHXUSPAIBRCJ	1930	Enigma I
	// II	AJDKSIRUXBLHWTMCQGZNPYFVOE	1930	Enigma I
	// III	BDFHJLCPRTXVZNYEIWGAKMUSQO	1930	Enigma I
	// IV	ESOVPZJAYQUIRHXLNFTGKDCMWB	December 1938	M3 Army
	// V	VZBRGITYUPSDNHLXAWMJQOFECK	December 1938	M3 Army
	// VI	JPGVOUMFYQBENHZRDKASXLICTW	1939	M3 & M4 Naval (FEB 1942)
	// VII	NZJHGRCXMYSWBOUFAIVLPEKQDT	1939	M3 & M4 Naval (FEB 1942)
	// VIII	FKQHTLXOCBJSPDZRAMEWNIUYGV	1939	M3 & M4 Naval (FEB 1942)
	var contacts []Contact

	codes := map[string]string{
		"I":   "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"II":  "AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"III": "BDFHJLCPRTXVZNYEIWGAKMUSQO",
		"IV":  "ESOVPZJAYQUIRHXLNFTGKDCMWB",
		"V":   "VZBRGITYUPSDNHLXAWMJQOFECK",
	}

	code := codes[number]
	for index, char := range alphabet {
		input := string(char)
		output := string(code[index])
		newContact := Contact{input: input, output: output}
		contacts = append(contacts, newContact)
	}

	return contacts

}

func CreateRotor(number string, startPosition int, ringSetting int) Rotor {
	// Rotor	Notch	Effect
	// I	Q	If rotor steps from Q to R, the next rotor is advanced
	// II	E	If rotor steps from E to F, the next rotor is advanced
	// III	V	If rotor steps from V to W, the next rotor is advanced
	// IV	J	If rotor steps from J to K, the next rotor is advanced
	// V	Z	If rotor steps from Z to A, the next rotor is advanced
	// VI, VII, VIII	Z+M	If rotor steps from Z to A, or from M to N the next rotor is advanced
	notches := map[string]string{
		"I":   "Q",
		"II":  "E",
		"III": "V",
		"IV":  "J",
		"V":   "Z",
	}

	return Rotor{
		position:      startPosition,
		ringSetting:   ringSetting,
		notchPosition: string(notches[number]),
		Contacts:      createContacts(number),
	}
}

func (r *Rotor) Rotate() {
	if r.position == maxPosition {
		r.position = minPosition
	} else {
		r.position++
	}
}

func (r *Rotor) GetOutput(input string) string {
	// Get the output for a given input
	for _, contact := range r.Contacts {
		if contact.input == input {
			return contact.output
		}
	}

	return ""
}

func (r *Rotor) GetInputPosition(input string) int {
	// get the position of a contact input value
	for position, contact := range r.Contacts {
		if contact.input == input {
			return position
		}
	}

	return 0
}

func (r *Rotor) GetOuputPosition(output string) int {
	// get the position of a contact output value
	for position, contact := range r.Contacts {
		if contact.output == output {
			return position
		}
	}

	return 0
}
