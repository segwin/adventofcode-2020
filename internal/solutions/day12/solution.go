package day12

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/segwin/adventofcode-2020/internal/geometry"
	"github.com/segwin/adventofcode-2020/internal/input"
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	s.part1(ctx, inputFile)
	s.part2(ctx, inputFile)
}

func (s *Solution) transformLines(scanner input.Scanner, ship, waypoint geometry.Point, part2 bool) (finalPosition geometry.Point, err error) {
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())
		if len(line) < 2 {
			return nil, fmt.Errorf("%w (%s)", ErrInvalidDirection, line)
		}

		directionCharacter := rune(line[0])
		magnitude, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		ship, waypoint, err = Transform(directionCharacter, int64(magnitude), ship, waypoint, part2)
		if err != nil {
			return nil, err
		}
	}

	return ship, nil
}

func (s *Solution) part1(ctx context.Context, inputFile string) {
	fmt.Println("\nPART 1")

	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to create input file scanner: %v\n", err)
		os.Exit(1)
	}

	initialShip := geometry.NewInts(0, 0)
	initialWaypoint := geometry.NewInts(1, 0) // start facing east

	finalShip, err := s.transformLines(scanner, initialShip, initialWaypoint, false)
	if err != nil {
		fmt.Printf("  ERROR: %v\n", err)
	}

	directionX := East
	if finalShip.MustGet(0).Int() < 0 {
		directionX = West
		finalShip.MustSet(geometry.Int(-finalShip.MustGet(0).Int()), 0)
	}

	directionY := North
	if finalShip.MustGet(1).Int() < 0 {
		directionY = South
		finalShip.MustSet(geometry.Int(-finalShip.MustGet(1).Int()), 1)
	}

	fmt.Printf("  RESULT: %s%d, %s%d\n", directionX, finalShip.MustGet(0).Int(), directionY, finalShip.MustGet(1).Int())
}

func (s *Solution) part2(ctx context.Context, inputFile string) {
	fmt.Println("\nPART 2")

	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to create input file scanner: %v\n", err)
		os.Exit(1)
	}

	initialShip := geometry.NewInts(0, 0)
	initialWaypoint := geometry.NewInts(10, 1) // start facing east

	finalShip, err := s.transformLines(scanner, initialShip, initialWaypoint, true)
	if err != nil {
		fmt.Printf("  ERROR: %v\n", err)
	}

	directionX := East
	if finalShip.MustGet(0).Int() < 0 {
		directionX = West
		finalShip.MustSet(geometry.Int(-finalShip.MustGet(0).Int()), 0)
	}

	directionY := North
	if finalShip.MustGet(1).Int() < 0 {
		directionY = South
		finalShip.MustSet(geometry.Int(-finalShip.MustGet(1).Int()), 1)
	}

	fmt.Printf("  RESULT: %s%d, %s%d\n", directionX, finalShip.MustGet(0).Int(), directionY, finalShip.MustGet(1).Int())
}
