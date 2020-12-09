package day1

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/segwin/advent-of-code/internal/input"
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	values, err := s.getValues(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to get values (%s)\n", err)
	}

	s.part1(values)
	s.part2(values)
}

func (s *Solution) getValues(ctx context.Context, inputFile string) (values []int, err error) {
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		return nil, err
	}

	defer scanner.Close()

	for scanner.Scan() {
		strValue := strings.TrimSpace(scanner.Text())

		value, err := strconv.Atoi(strValue)
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	return values, nil
}

func (s *Solution) part1(values []int) {
	fmt.Println("PART 1")

	for i := range values {
		for j := i + 1; j < len(values); j++ {
			if values[i]+values[j] == 2020 {
				fmt.Printf("  Found %d (%d) + %d (%d) = 2020\n", values[i], i, values[j], j)
				fmt.Printf("  RESULT: %d*%d = %d\n", values[i], values[j], values[i]*values[j])
			}
		}
	}
}

func (s *Solution) part2(values []int) {
	fmt.Println("PART 2")

	for i := range values {
		for j := i + 1; j < len(values); j++ {
			for k := j + 1; k < len(values); k++ {
				if values[i]+values[j]+values[k] == 2020 {
					fmt.Printf("  Found %d (%d) + %d (%d) + %d (%d) = 2020\n", values[i], i, values[j], j, values[k], k)
					fmt.Printf("  RESULT: %d*%d+%d = %d\n", values[i], values[j], values[k], values[i]*values[j]*values[k])
				}
			}
		}
	}
}
