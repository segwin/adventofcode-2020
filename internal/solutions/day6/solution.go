package day6

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/segwin/advent-of-code/internal/input"
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	part1Scanner := s.NewScanner(ctx, inputFile)
	defer part1Scanner.Close()

	s.part1(part1Scanner)

	part2Scanner := s.NewScanner(ctx, inputFile)
	defer part2Scanner.Close()

	s.part2(part2Scanner)
}

func (s *Solution) NewScanner(ctx context.Context, inputFile string) input.Scanner {
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to create scanner: %v", err)
		os.Exit(1)
	}

	return scanner
}

func (s *Solution) getLines(scanner input.Scanner, callback func(line string)) {
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("File read error: %v\n", err)
			os.Exit(1)
		}

		line := strings.TrimSpace(scanner.Text())
		callback(line)
	}
}

func (s *Solution) part1(scanner input.Scanner) {
	fmt.Println("PART 1")

	yesCount := 0
	group := NewIndividualResponses()

	callback := func(line string) {
		if len(line) == 0 {
			// hit end of group, tally up the group's responses & reset for next one
			yesCount += group.YesCount()
			group = NewIndividualResponses()

			return
		}

		if err := group.UnmarshalNew(line); err != nil {
			fmt.Printf("  Failed to unmarshal response: %v (line = %q)\n", err, line)
			os.Exit(1)
		}
	}

	s.getLines(scanner, callback)

	// handle case where file ends without an additional newline (operation is safe even if currentGroup == nil)
	yesCount += group.YesCount()

	// print result
	fmt.Printf("  RESULT: Got %d affirmatives across all groups\n", yesCount)
}

func (s *Solution) part2(scanner input.Scanner) {
	fmt.Println("PART 2")

	yesCount := 0
	group := NewUnanimousResponses()

	callback := func(line string) {
		if len(line) == 0 {
			// hit end of group, tally up the group's responses & reset for next one
			yesCount += group.YesCount()
			group = NewUnanimousResponses()

			return
		}

		if err := group.UnmarshalNew(line); err != nil {
			fmt.Printf("  Failed to unmarshal response for: %v (line = %q)\n", err, line)
			os.Exit(1)
		}
	}

	s.getLines(scanner, callback)

	// handle case where file ends without an additional newline (operation is safe even if currentGroup == nil)
	yesCount += group.YesCount()

	// print result
	fmt.Printf("  RESULT: Got %d affirmatives across all groups\n", yesCount)
}
