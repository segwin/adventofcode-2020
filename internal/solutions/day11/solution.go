package day11

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

	layout, err := s.getLayout(scanner)
	if err != nil {
		fmt.Printf("ERROR: Failed to get layout: %v\n", err)
		os.Exit(1)
	}

	s.part1(layout)
	s.part2(layout)
}

func (s *Solution) getLayout(scanner input.Scanner) (layout Layout, err error) {
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())
		row, err := ParseRow(line)
		if err != nil {
			return nil, err
		}

		layout = append(layout, row)
	}

	return layout, nil
}

func (s *Solution) part1(layout Layout) {
	fmt.Println("\nPART 1")

	prevLayout := layout.Clone()
	generation := 0
	for {
		newLayout := s.evolveLayout(prevLayout, 4, false)
		if newLayout.Equals(prevLayout) {
			// reached equilibrium
			fmt.Printf("  RESULT: Found %d occupied seats (evolution took %d generations)\n", newLayout.Count(Occupied), generation)
			return
		}

		prevLayout = newLayout
		generation++
	}
}

func (s *Solution) part2(layout Layout) {
	fmt.Println("\nPART 2")

	prevLayout := layout.Clone()
	generation := 0
	for {
		newLayout := s.evolveLayout(prevLayout, 5, true)
		if newLayout.Equals(prevLayout) {
			// reached equilibrium
			fmt.Printf("  RESULT: Found %d occupied seats (evolution took %d generations)\n", newLayout.Count(Occupied), generation)
			return
		}

		prevLayout = newLayout
		generation++
	}
}

func (s *Solution) evolveLayout(prevLayout Layout, tolerateOccupied int, skipAdjacentFloors bool) (newLayout Layout) {
	newLayout = prevLayout.Clone()

	for row := range prevLayout {
		for col := range prevLayout[row] {
			adjacentSeats, err := prevLayout.AdjacentSeats(row, col, skipAdjacentFloors)
			if err != nil {
				fmt.Printf("  ERROR: Failed to get adjacent seats: %s\n", err)
			}

			occupiedAdjacent := prevLayout.FilterByType(adjacentSeats, Occupied)

			switch prevLayout[row][col] {
			case Empty:
				if len(occupiedAdjacent) == 0 {
					// this seat becomes occupied
					newLayout[row][col] = Occupied
				}

			case Occupied:
				if len(occupiedAdjacent) >= tolerateOccupied {
					// this seat becomes empty
					newLayout[row][col] = Empty
				}

			case Floor:
				// no action
			}
		}
	}

	return newLayout
}
