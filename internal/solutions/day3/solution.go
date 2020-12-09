package day3

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/segwin/advent-of-code/internal/input"
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	slopes := []Position{
		{X: 1, Y: 1},
		{X: 3, Y: 1}, // part 1
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}

	navMap, err := s.getMap(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to get map from input: %s\n", err)
		os.Exit(1)
	}

	s.run(1, navMap, slopes[1])
	s.run(2, navMap, slopes...)
}

func (s *Solution) getMap(ctx context.Context, inputFile string) (navMap *Map, err error) {
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		return nil, err
	}

	defer scanner.Close()

	var rows []Row
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		var row Row
		if err := row.UnmarshalPattern(line); err != nil {
			return nil, err
		}

		rows = append(rows, row)
	}

	return &Map{Rows: rows}, nil
}

func (s *Solution) run(part int, navMap *Map, slopes ...Position) {
	fmt.Printf("\nPART %d\n", part)

	hitsProduct := int64(1)
	for _, slope := range slopes {
		hits := navMap.CountHits(slope.X, slope.Y)
		hitsProduct *= int64(hits)

		fmt.Printf("  With slope %+v, hit %d trees\n", slope, hits)
	}

	fmt.Printf("  RESULT: Product of all hits is %d\n", hitsProduct)
}
