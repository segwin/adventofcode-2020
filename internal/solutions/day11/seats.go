package day11

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidSeatCharacter = errors.New("invalid seat character")
	ErrInvalidSeatPosition  = errors.New("invalid seat position")
)

type Seat rune

const (
	Floor    Seat = '.'
	Empty    Seat = 'L'
	Occupied Seat = '#'
)

func (s Seat) String() string {
	switch s {
	case Floor:
		return "floor"
	case Empty:
		return "empty"
	case Occupied:
		return "occupied"
	}

	return "invalid"
}

func ParseRow(encoded string) (row []Seat, err error) {
	for _, s := range encoded {
		switch seat := Seat(s); seat {
		case Floor, Empty, Occupied:
			row = append(row, seat)
		default:
			return nil, fmt.Errorf("%w (%s)", ErrInvalidSeatCharacter, string(s))
		}
	}

	return row, nil
}

type Layout [][]Seat

func (l Layout) Clone() Layout {
	cloned := make([][]Seat, len(l))
	for i, row := range l {
		cloned[i] = make([]Seat, len(row))
		copy(cloned[i], row)
	}

	return cloned
}

func (l Layout) Equals(other Layout) bool {
	if len(l) != len(other) {
		return false
	}

	for row := range l {
		if len(l[row]) != len(other[row]) {
			return false
		}

		for col := range l[row] {
			if l[row][col] != other[row][col] {
				return false
			}
		}
	}

	return true
}

func (l Layout) AdjacentSeats(row, col int, skipFloor bool) (adjacentSeats []Seat, err error) {
	if row > len(l)-1 || col > len(l[row])-1 {
		return nil, fmt.Errorf("%w (%d:%d)", ErrInvalidSeatPosition, row, col)
	}

	directions := [][]int{
		{0, 1},   // right
		{1, 1},   // lower-right
		{1, 0},   // down
		{1, -1},  // lower-left
		{0, -1},  // left
		{-1, -1}, // upper-left
		{-1, 0},  // up
		{-1, 1},  // upper-right
	}

	for _, direction := range directions {
		if seat := l.nextInDirection(row, col, direction[0], direction[1], skipFloor); seat != nil {
			adjacentSeats = append(adjacentSeats, *seat)
		}
	}

	return adjacentSeats, nil
}

func (l Layout) nextInDirection(row, col, rowStep, colStep int, skipFloor bool) (seat *Seat) {
	for l.withinBounds(row+rowStep, col+colStep) {
		row += rowStep
		col += colStep

		seat = &l[row][col]

		if skipFloor {
			if *seat == Floor {
				continue // try next in this direction
			}
		}

		// not skipping floor spaces: only one iteration allowed
		return seat
	}

	return nil
}

func (l Layout) withinBounds(row, col int) bool {
	if row < 0 || row > len(l)-1 {
		return false
	}

	if col < 0 || col > len(l[row])-1 {
		return false
	}

	return true
}

func (l Layout) FilterByType(seats []Seat, allowedTypes ...Seat) (filtered []Seat) {
	for _, seat := range seats {
		for _, seatType := range allowedTypes {
			if seat == seatType {
				filtered = append(filtered, seat)
				break
			}
		}
	}

	return filtered
}

func (l Layout) Count(seatType Seat) (count int) {
	for row := range l {
		for col := range l[row] {
			if l[row][col] == seatType {
				count++
			}
		}
	}

	return count
}

func (l Layout) String() (layoutStr string) {
	for row := range l {
		for col := range l[row] {
			layoutStr += string(rune(l[row][col]))
		}
		layoutStr += "\n"
	}

	return layoutStr
}
