package day5

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/segwin/advent-of-code/internal/input"
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	seatIDs, err := s.getSeats(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to parse seats: %s\n", err)
		os.Exit(1)
	}

	s.part1(seatIDs...)
	s.part2(seatIDs...)
}

func (s *Solution) getSeats(ctx context.Context, inputFile string) (seatIDs []uint16, err error) {
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		return nil, err
	}

	defer scanner.Close()

	// read all lines, keeping track of the highest seat ID along the way
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		encodedSeat := strings.TrimSpace(scanner.Text())

		var seat Seat
		if err := seat.Unmarshal(encodedSeat); err != nil {
			return nil, err
		}

		seatIDs = append(seatIDs, seat.ID())
	}

	return seatIDs, nil
}

func (s *Solution) part1(seatIDs ...uint16) {
	fmt.Println("PART 1")

	maxID := uint16(0)
	for _, id := range seatIDs {
		if id > maxID {
			maxID = id
		}
	}

	fmt.Printf("  RESULT: Highest seat ID is %d\n", maxID)
}

func (s *Solution) part2(seatIDs ...uint16) {
	fmt.Println("PART 2")

	// sort seat IDs so we can identify a gap
	sort.Slice(seatIDs, func(i, j int) bool { return seatIDs[i] < seatIDs[j] })

	myID := uint16(0)
	for i := 0; i < len(seatIDs)-1; i++ {
		if diff := seatIDs[i+1] - seatIDs[i]; diff > 1 {
			myID = seatIDs[i] + 1
			break // no other gaps are allowed by the problem statement
		}
	}

	fmt.Printf("  RESULT: My seat ID is %d\n", myID)
}
