package day9

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/segwin/adventofcode-2020/internal/input"
)

var (
	ErrNonXMASValueNotFound = errors.New("failed to find non-XMAS-compliant value")
	ErrWeaknessNotFound     = errors.New("failed to find encryption weakness")
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	values, err := s.getValues(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to get values: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("\nPART 1")
	invalidValue, err := s.part1(values)
	if err != nil {
		fmt.Printf("  ERROR: %s\n", err)
	}

	fmt.Printf("  RESULT: The first value that doesn't follow XMAS is %d\n", invalidValue)

	fmt.Println("\nPART 2")
	weakness, err := s.part2(values, invalidValue)
	if err != nil {
		fmt.Printf("  ERROR: %s\n", err)
	}

	fmt.Printf("  RESULT: The encryption weakness is %d\n", weakness)
}

func (s *Solution) getValues(ctx context.Context, inputFile string) (values []int, err error) {
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		return nil, err
	}

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	return values, nil
}

func (s *Solution) part1(values []int) (invalidValue int, err error) {
	last25 := NewIntRing(25)
	for _, value := range values {
		if last25.Len() == 25 {
			if !s.followsXMAS(value, last25) {
				return value, nil
			}
		}

		// add to last25 & continue with next line
		if err := last25.Push(value); err != nil {
			return 0, err
		}
	}

	return 0, ErrNonXMASValueNotFound
}

func (s *Solution) part2(values []int, invalidValue int) (weakness int, err error) {
	for offset := range values {
		sum := 0
		var positions []int

		reset := func() {
			sum = 0
			positions = nil
		}

		for i := offset; i < len(values); i++ {
			// special case: skip the invalid value itself
			if values[i] == invalidValue {
				reset()
				continue
			}

			sum += values[i]
			positions = append(positions, i)

			if sum == invalidValue {
				min := math.MaxInt32
				max := math.MinInt32

				for _, position := range positions {
					if values[position] < min {
						min = values[position]
					}
					if values[position] > max {
						max = values[position]
					}
				}

				return min + max, nil
			}

			if sum > invalidValue {
				reset() // this sum didn't work out, reset & try with next position
			}
		}
	}

	return 0, ErrWeaknessNotFound
}

func (s *Solution) followsXMAS(value int, last25 Ring) bool {
	for i := 0; i < last25.Len(); i++ {
		last1 := s.getAt(i, last25)

		for j := 0; j < last25.Len(); j++ {
			if j == i {
				continue
			}

			last2 := s.getAt(j, last25)
			if last1+last2 == value {
				// ok
				return true
			}
		}
	}

	return false
}

func (s *Solution) getAt(i int, last25 Ring) int {
	value, ok := last25.MustGet(i).(int)
	if !ok {
		fmt.Printf("ERROR: Got non-integer value in last25 ring")
		os.Exit(1)
	}

	return value
}
