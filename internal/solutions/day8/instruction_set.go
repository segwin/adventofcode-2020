package day8

type instructionSet []*instruction

// Execute this instruction set.
func (s instructionSet) Execute() (accumulator int, sequence []int, recursionDetected bool) {
	for i := 0; i < len(s); {
		if s[i].hit {
			recursionDetected = true
			break // repeating a previously executed instruction
		}

		sequence = append(sequence, i)
		s[i].hit = true

		switch s[i].operation {
		case jmp:
			i += s[i].argument
			continue // go to new instruction immediately

		case nop:
			// do nothing

		case acc:
			accumulator += s[i].argument
		}

		i++
	}

	return accumulator, sequence, recursionDetected
}

// Reset hit indicators for all instructions.
func (s instructionSet) Reset() {
	for _, instruction := range s {
		instruction.hit = false
	}
}
