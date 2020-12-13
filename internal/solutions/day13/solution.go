package day13

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/segwin/adventofcode-2020/internal/input"
)

var (
	ErrInvalidNotes = errors.New("invalid input notes")
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

	earliestDeparture, busList, err := s.parsePart1Notes(scanner)
	if err != nil {
		fmt.Printf("  ERROR: Failed to get problem values: %v\n", err)
		os.Exit(1)
	}

	minWaitTime := time.Duration(math.MaxInt64)
	var fastestBus *Bus

	for _, bus := range busList.Buses {
		waitTime := bus.DepartsIn(earliestDeparture)
		if waitTime < minWaitTime {
			fastestBus = bus
			minWaitTime = waitTime
		}
	}

	waitMinutes := int(minWaitTime.Minutes())
	fmt.Printf("  Bus %d is the fastest with a %d minute wait time\n", fastestBus.ID, waitMinutes)
	fmt.Printf("  RESULT: ID * wait minutes => %d\n", fastestBus.ID*waitMinutes)
}

func (s *Solution) parsePart1Notes(scanner input.Scanner) (earliestDeparture time.Duration, busList *BusList, err error) {
	lines := make([]string, 0, 2)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, nil, err
		}

		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	if len(lines) != 2 {
		return 0, nil, fmt.Errorf("%w (got %d lines, expected 2)", ErrInvalidNotes, len(lines))
	}

	estimatedMinutes, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, nil, err
	}

	busList = &BusList{}
	if err := busList.Unmarshal(lines[1]); err != nil {
		return 0, nil, err
	}

	return time.Minute * time.Duration(estimatedMinutes), busList, nil
}

func (s *Solution) part2(ctx context.Context, inputFile string) {
	fmt.Println("\nPART 2")

	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		fmt.Printf("  ERROR: Failed to create input file scanner: %v\n", err)
		os.Exit(1)
	}

	schedule, err := s.parsePart2Notes(scanner)
	if err != nil {
		fmt.Printf("  ERROR: Failed to get problem values: %v\n", err)
		os.Exit(1)
	}

	lowestTime, err := schedule.FindLowestTime()
	if err != nil {
		fmt.Printf("  ERROR: Failed to compute lowest time: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("  RESULT: All buses coincide with the schedule in %d minutes\n", lowestTime)
}

func (s *Solution) parsePart2Notes(scanner input.Scanner) (schedule *Schedule, err error) {
	lines := make([]string, 0, 2)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	if len(lines) != 2 {
		return nil, fmt.Errorf("%w (got %d lines, expected 2)", ErrInvalidNotes, len(lines))
	}

	schedule = &Schedule{}
	if err := schedule.Unmarshal(lines[1]); err != nil {
		return nil, err
	}

	return schedule, nil
}
