package day14

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/segwin/adventofcode-2020/internal/input"
)

var (
	ErrUnexpectedPermutations = errors.New("unexpected number of permutations")
	ErrBadInputLine           = errors.New("invalid input line")
	ErrBadMemFormat           = errors.New("invalid mem[N] format")
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	s.part1(ctx, inputFile)
	s.part2(ctx, inputFile)
}

func (s *Solution) part1(ctx context.Context, inputFile string) {
	fmt.Println("\nPART 1")

	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		fmt.Printf("  ERROR: Failed to create input file scanner: %v\n", err)
		os.Exit(1)
	}

	memory, err := s.runLines(scanner, true)
	if err != nil {
		fmt.Printf("  ERROR: Failed to get problem values: %v\n", err)
		os.Exit(1)
	}

	// sum all values in memory
	sum := int64(0)
	for _, value := range memory {
		sum += value.Int()
	}

	fmt.Printf("  RESULT: The sum of all stored values is %d\n", sum)
}

func (s *Solution) part2(ctx context.Context, inputFile string) {
	fmt.Println("\nPART 2")

	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		fmt.Printf("  ERROR: Failed to create input file scanner: %v\n", err)
		os.Exit(1)
	}

	memory, err := s.runLines(scanner, false)
	if err != nil {
		fmt.Printf("  ERROR: Failed to get problem values: %v\n", err)
		os.Exit(1)
	}

	// sum all values in memory
	sum := int64(0)
	for _, value := range memory {
		sum += value.Int()
	}

	fmt.Printf("  RESULT: The sum of all stored values is %d\n", sum)
}

func (s *Solution) runLines(scanner input.Scanner, part1 bool) (memory map[int64]*Bitset, err error) {
	memory = map[int64]*Bitset{}

	currentMask := Bitmask{}
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())

		// expect format: "<operation> = <value>"
		words := strings.Split(line, " = ")
		if len(words) != 2 {
			return nil, fmt.Errorf("%w (%q)", ErrBadInputLine, line)
		}

		if words[0] == "mask" {
			// simpler case: just update the mask
			if err := currentMask.Unmarshal(words[1], part1); err != nil {
				return nil, err
			}

			continue
		}

		// case 2: memory
		address, value, err := s.parseMem(words[0], words[1])
		if err != nil {
			return nil, err
		}

		if part1 {
			maskedValues, err := currentMask.Apply(value)
			if err != nil {
				return nil, err
			}

			// len(maskedValues) can only ever be 1
			if len(maskedValues) != 1 {
				return nil, ErrUnexpectedPermutations
			}

			memory[address.Int()] = maskedValues[0]
		} else {
			maskedAddresses, err := currentMask.Apply(address)
			if err != nil {
				return nil, err
			}

			for _, maskedAddress := range maskedAddresses {
				memory[maskedAddress.Int()] = value
			}
		}
	}

	return memory, nil
}

func (s *Solution) parseMem(key, value string) (address *Bitset, bitset *Bitset, err error) {
	matches := regexp.MustCompile(`mem\[(?P<address>[0-9]+)\]`).FindStringSubmatch(key)
	if len(matches) != 2 {
		return nil, nil, fmt.Errorf("%w (%q)", ErrBadMemFormat, key)
	}

	addressStr := matches[1]
	address = &Bitset{}
	if err := address.Unmarshal(addressStr, 36); err != nil {
		return nil, nil, err
	}

	bitset = &Bitset{}
	if err := bitset.Unmarshal(value, 36); err != nil {
		return nil, nil, err
	}

	return address, bitset, nil
}
