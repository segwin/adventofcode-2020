package day8

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/segwin/advent-of-code/internal/input"
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	instructions, err := s.getInstructions(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to parse instructions: %s\n", err)
		os.Exit(1)
	}

	part1Sequence := part1(instructions)
	part2(instructions, part1Sequence)
}

func (s *Solution) getInstructions(ctx context.Context, inputFile string) (instructions instructionSet, err error) {
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		return nil, err
	}

	defer scanner.Close()

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())

		var newInstruction instruction
		if err := newInstruction.Unmarshal(line); err != nil {
			return nil, err
		}

		instructions = append(instructions, &newInstruction)
	}

	return instructions, nil
}

func part1(instructions instructionSet) (sequence []int) {
	accumulator, sequence, recursionDetected := instructions.Execute()
	if recursionDetected {
		fmt.Printf("PART 1 NOTE: detected recursion (last position: %d)\n", sequence[len(sequence)-1])
	}

	fmt.Printf("PART 1 RESULT: Accumulator = %d\n", accumulator)

	return sequence
}

func tryFixAt(instructions instructionSet, fixPosition int) (ok bool) {
	instruction := instructions[fixPosition]

	previousOp := instruction.operation
	defer func() {
		// revert fix before returning in case this isn't "the one"
		instruction.operation = previousOp
	}()

	switch instruction.operation {
	case jmp:
		instruction.operation = nop

	case nop:
		if instruction.argument == 1 {
			return false // nop +1 and jmp +1 are equivalent here since both jump exactly 1 instruction forward
		}
		instruction.operation = jmp

	case acc:
		return false // acc can't change the outcome here
	}

	// get result
	accumulator, _, recursionDetected := instructions.Execute()
	if recursionDetected {
		return false // try again with the previous instruction in the sequence
	}

	fmt.Printf("PART 2 RESULT: Accumulator = %d\n", accumulator)
	return true
}

func part2(instructions instructionSet, part1Sequence []int) {
	// start at the end of the previous sequence and try flipping nop<->jmp until a working path is found
	for s := len(part1Sequence) - 1; s >= 0; s-- {
		instructions.Reset()
		if ok := tryFixAt(instructions, part1Sequence[s]); ok {
			return
		}
	}

	fmt.Printf("PART 2 ERROR: Failed to find any instructions where flipping jmp<->nop removed the infinite recursion\n")
	os.Exit(1)
}