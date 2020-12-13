package day19

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/segwin/adventofcode-2020/internal/input"
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to create input file scanner: %v\n", err)
		os.Exit(1)
	}

	values, err := s.getValues(scanner)
	if err != nil {
		fmt.Printf("ERROR: Failed to get problem values: %v\n", err)
		os.Exit(1)
	}

	s.part1(values)
	s.part2(values)
}

func (s *Solution) getValues(scanner input.Scanner) (values []ProblemValue, err error) {
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())

		var v ProblemValue
		if err := v.Unmarshal(line); err != nil {
			return nil, err
		}

		values = append(values, v)
	}

	return values, nil
}

func (s *Solution) part1(values []ProblemValue) {
	fmt.Println("\nPART 1")

	// TODO
}

func (s *Solution) part2(values []ProblemValue) {
	fmt.Println("\nPART 2")

	// TODO
}
