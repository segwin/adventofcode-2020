package day15

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/segwin/adventofcode-2020/internal/input"
)

var (
	ErrInvalidInput = errors.New("invalid input file")
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to create input file scanner: %v\n", err)
		os.Exit(1)
	}

	startingNumbers, err := s.readInput(scanner)
	if err != nil {
		fmt.Printf("ERROR: Failed to get input line: %v\n", err)
		os.Exit(1)
	}

	memory := NewMemory()
	if err := memory.Unmarshal(startingNumbers); err != nil {
		fmt.Printf("  ERROR: Failed to parse starting numbers: %v\n", err)
		return
	}

	lastSpoken := s.play(memory, 1, 2020, 0)
	s.play(memory, 2, 30000000, lastSpoken)
}

func (s *Solution) readInput(scanner input.Scanner) (line string, err error) {
	var lines []string
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}

		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	if len(lines) != 1 {
		return "", fmt.Errorf("%w (got %d lines)", ErrInvalidInput, len(lines))
	}

	return lines[0], nil
}

func (s *Solution) play(memory Memory, part int, generations int, lastSpoken int) (newLastSpoken int) {
	fmt.Printf("\nPART %d\n", part)

	lastSpoken = memory.Recite(generations, lastSpoken)

	fmt.Printf("  RESULT: Number spoken on %dth iteration is %d\n", generations, lastSpoken)
	return lastSpoken
}
