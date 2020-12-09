package day4

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/segwin/advent-of-code/internal/input"
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	passports, err := s.getPassports(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to parse passports: %s\n", err)
		os.Exit(1)
	}

	s.run(1, passports, false)
	s.run(2, passports, true)
}

// parsePassport reads all lines in the input file, collecting passport lines along
// the way & unmarshaling when end-of-passport is reached.
func (s *Solution) getPassports(ctx context.Context, inputFile string) (passports []Passport, err error) {
	// create scanner
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		return nil, err
	}

	defer scanner.Close()

	// parse all passports
	parseAndAppend := func(lines []string) error {
		passport := Passport{}
		if err := passport.Unmarshal(lines); err != nil {
			return err
		}

		passports = append(passports, passport)
		return nil
	}

	var passportLines []string
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("File read error: %v\n", err)
			os.Exit(1)
		}

		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			passportLines = append(passportLines, line)
			continue
		}

		// reached end of passport
		if err := parseAndAppend(passportLines); err != nil {
			return nil, err
		}

		passportLines = nil // reset for next passport
	}

	if len(passportLines) > 0 {
		if err := parseAndAppend(passportLines); err != nil {
			return nil, err
		}
	}

	return passports, nil
}

func (s *Solution) run(part int, passports []Passport, checkValues bool) {
	fmt.Printf("\nPART %d\n", part)

	validCount := 0
	for _, passport := range passports {
		if passport.IsValid(checkValues) {
			validCount++
		}
	}

	fmt.Printf("  RESULT: Found %d valid passports\n", validCount)
}
