package day3

import (
	"errors"
	"fmt"
)

var (
	ErrUnknownSymbol = errors.New("unknown map symbol")
)

type Symbol rune

const (
	Tree Symbol = '#'
	Open Symbol = '.'
)

type Row struct {
	pattern []Symbol
}

func (r *Row) UnmarshalPattern(line string) error {
	for _, s := range line {
		switch symbol := Symbol(s); symbol {
		case Tree, Open:
			r.pattern = append(r.pattern, symbol)
		default:
			return fmt.Errorf("%w (%s)", ErrUnknownSymbol, string(symbol))
		}
	}

	return nil
}

func (r *Row) IsOpen(position int) bool {
	return r.pattern[position%len(r.pattern)] == Open
}

type Position struct {
	X int
	Y int
}

func (p *Position) Add(x, y int) {
	p.X += x
	p.Y += y
}

type Map struct {
	Rows []Row
}

func (m *Map) CountHits(rightStep, downStep int) (count int) {
	position := Position{}

	for position.Y < len(m.Rows) {
		if !m.Rows[position.Y].IsOpen(position.X) {
			count++
		}

		position.Add(rightStep, downStep)
	}

	return count
}
