package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

func main() {
	f, err := ioutil.ReadFile("day3/input")
	if err != nil {
		fmt.Printf("Read error\n")
		panic(err)
	}

	lines := strings.Split(strings.Trim(string(f), " \t\v\n"), "\n")
	navMap := Map{Rows: make([]Row, len(lines))}

	for i, line := range lines {
		if err := navMap.Rows[i].UnmarshalPattern(line); err != nil {
			fmt.Printf("Fatal error: %s\n", err)
			os.Exit(1)
		}
	}

	slopes := []Position{
		{X: 1, Y: 1},
		{X: 3, Y: 1}, // part 1
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}

	hitsProduct := int64(1)
	for _, slope := range slopes {
		hits := navMap.CountHits(slope.X, slope.Y)
		hitsProduct *= int64(hits)

		fmt.Printf("INTERMEDIATE: With slope %+v, hit %d trees\n", slope, hits)
	}

	fmt.Printf("RESULT: Product of all hits is %d\n", hitsProduct)
}
