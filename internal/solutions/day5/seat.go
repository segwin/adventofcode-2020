package day5

import (
	"errors"
	"fmt"
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
