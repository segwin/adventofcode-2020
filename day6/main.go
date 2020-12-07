package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func getLines(scanner *bufio.Scanner, callback func(line string)) {
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("File read error: %v\n", err)
			os.Exit(1)
		}

		line := strings.Trim(scanner.Text(), " \t\v\n")
		callback(line)
	}
}

func part1(scanner *bufio.Scanner) {
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
			fmt.Printf("Failed to unmarshal response for line %q: %v\n", line, err)
			os.Exit(1)
		}
	}

	getLines(scanner, callback)

	// handle case where file ends without an additional newline (operation is safe even if currentGroup == nil)
	yesCount += group.YesCount()

	// print result
	fmt.Printf("PART 1 RESULT: Got %d affirmatives across all groups\n", yesCount)
}

func part2(scanner *bufio.Scanner) {
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
			fmt.Printf("Failed to unmarshal response for line %q: %v\n", line, err)
			os.Exit(1)
		}
	}

	getLines(scanner, callback)

	// handle case where file ends without an additional newline (operation is safe even if currentGroup == nil)
	yesCount += group.YesCount()

	// print result
	fmt.Printf("PART 2 RESULT: Got %d affirmatives across all groups\n", yesCount)
}

func main() {
	file, err := os.Open("day6/input")
	if err != nil {
		fmt.Printf("File open error: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	part1(bufio.NewScanner(file))

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		fmt.Printf("Failed to reset scanner position: %v\n", err)
		os.Exit(1)
	}

	part2(bufio.NewScanner(file))
}
