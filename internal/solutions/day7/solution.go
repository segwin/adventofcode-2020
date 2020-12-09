package day7

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/segwin/advent-of-code/internal/input"
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	bags, err := s.getBags(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to parse bags: %s\n", err)
		os.Exit(1)
	}

	part1("shiny gold", bags)
	part2("shiny gold", bags)
}

func (s *Solution) getBags(ctx context.Context, inputFile string) (bags map[string]Bag, err error) {
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		return nil, err
	}

	defer scanner.Close()

	bags = map[string]Bag{}
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())

		newBag := &bag{}
		if err := newBag.Unmarshal(line); err != nil {
			return nil, err
		}

		bags[newBag.Colour()] = newBag
	}

	return bags, nil
}

func part1(colour string, bags map[string]Bag) {
	fmt.Println("PART 1")

	count := 0
	for _, bag := range bags {
		if bag.Contains(colour, bags) {
			count++
		}
	}

	// print result
	fmt.Printf("  RESULT: Found %d bags that can contain a %s bag\n", count, colour)
}

func part2(colour string, bags map[string]Bag) {
	fmt.Println("PART 2")

	bag, ok := bags[colour]
	if !ok {
		fmt.Printf("  ERROR: Failed to find %s bag\n", colour)
		os.Exit(1)
	}

	count := int64(0)
	subBags := bag.SubBags(bags)
	for _, subBagCount := range subBags {
		count += subBagCount
	}

	// print result
	fmt.Printf("  RESULT: Found total of %d bags inside the %s bag\n", count, colour)
}
