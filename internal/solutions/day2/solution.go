package day2

import (
	"context"
	"fmt"
	"strings"

	"github.com/segwin/advent-of-code/internal/input"
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	s.part1(ctx, inputFile)
	s.part2(ctx, inputFile)
}

func (s *Solution) getEntries(ctx context.Context, inputFile string, oldPolicy bool) (entries []*PasswordEntry, err error) {
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		return nil, err
	}

	defer scanner.Close()

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		entry, err := UnmarshalEntry(line, oldPolicy)
		if err != nil {
			fmt.Printf("Parse error: %s (line: %q)\n", err, line)
			continue
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func (s *Solution) part1(ctx context.Context, inputFile string) {
	fmt.Println("PART 1")

	entries, err := s.getEntries(ctx, inputFile, true)
	if err != nil {
		fmt.Printf("  ERROR: Failed to get values using old password policy (%s)\n", err)
		return
	}

	validCount := 0
	for _, entry := range entries {
		if entry.IsValid() {
			validCount++
		}
	}

	fmt.Printf("  RESULT: Got %d valid lines\n", validCount)
}

func (s *Solution) part2(ctx context.Context, inputFile string) {
	fmt.Println("PART 2")

	entries, err := s.getEntries(ctx, inputFile, false)
	if err != nil {
		fmt.Printf("ERROR: Failed to get values using current password policy (%s)\n", err)
		return
	}

	validCount := 0
	for _, entry := range entries {
		if entry.IsValid() {
			validCount++
		}
	}

	fmt.Printf("  RESULT: Got %d valid lines\n", validCount)
}
