package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	ErrInvalidEncodedLength = errors.New("invalid encoded seat ID length")
	ErrInvalidEncodingChar  = errors.New("invalid seat ID encoding character")
)

type Seat struct {
	Row    uint8 // 7 last bits used
	Column uint8 // 3 last bits used
}

func (s *Seat) Unmarshal(encoded string) (err error) {
	if len(encoded) != 10 {
		return fmt.Errorf("%w (%q)", ErrInvalidEncodedLength, encoded)
	}

	row, err := s.toByteMask(encoded[:7], 'F', 'B')
	if err != nil {
		return err
	}

	column, err := s.toByteMask(encoded[7:], 'L', 'R')
	if err != nil {
		return err
	}

	// ok
	s.Row = row
	s.Column = column

	return nil
}

func (s *Seat) ID() uint16 {
	return (uint16(s.Row) << 3) | uint16(s.Column)
}

func (s *Seat) toByteMask(id string, low rune, high rune) (mask uint8, err error) {
	highestBit := len(id) - 1
	for i, value := range id {
		var bit uint8
		switch value {
		case low:
			bit = 0
		case high:
			bit = 1
		default:
			return 0, fmt.Errorf("%w (%s)", ErrInvalidEncodingChar, string(value))
		}

		mask |= (bit << (highestBit - i))
	}

	return mask, nil
}

func part1(seatIDs ...uint16) {
	maxID := uint16(0)
	for _, id := range seatIDs {
		if id > maxID {
			maxID = id
		}
	}

	fmt.Printf("PART 1 RESULT: Highest seat ID is %d\n", maxID)
}

func part2(seatIDs ...uint16) {
	// sort seat IDs so we can identify a gap
	sort.Slice(seatIDs, func(i, j int) bool { return seatIDs[i] < seatIDs[j] })

	myID := uint16(0)
	for i := 0; i < len(seatIDs)-1; i++ {
		if diff := seatIDs[i+1] - seatIDs[i]; diff > 1 {
			myID = seatIDs[i] + 1
			break // no other gaps are allowed by the problem statement
		}
	}

	fmt.Printf("PART 2 RESULT: My seat ID is %d\n", myID)
}

func main() {
	file, err := os.Open("day5/input")
	if err != nil {
		fmt.Printf("File open error: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read all lines, keeping track of the highest seat ID along the way
	var seatIDs []uint16

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("File read error: %v\n", err)
			os.Exit(1)
		}

		encodedSeat := strings.Trim(scanner.Text(), " \t\v\n")

		var seat Seat
		if err := seat.Unmarshal(encodedSeat); err != nil {
			fmt.Printf("Failed to unmarshal seat ID: %v", err)
			os.Exit(1)
		}

		seatIDs = append(seatIDs, seat.ID())
	}

	part1(seatIDs...)
	part2(seatIDs...)
}
