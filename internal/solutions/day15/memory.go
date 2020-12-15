package day15

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrNoStartingNumbers = errors.New("no starting numbers in memory")
)

type Memory interface {
	Unmarshal(startingNumbers string) error
	Recite(untilGeneration int, lastSpoken int) (newLastSpoken int)
}

type memory struct {
	StartingNumbers []int

	previousNumbers map[int]int // maps spoken numbers -> generation spoken
	generation      int
}

func NewMemory() Memory {
	return &memory{
		previousNumbers: map[int]int{},
		generation:      0,
	}
}

func (m *memory) Unmarshal(startingNumbers string) error {
	for _, numberStr := range strings.Split(startingNumbers, ",") {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			return err
		}

		m.StartingNumbers = append(m.StartingNumbers, number)
	}

	return nil
}

func (m *memory) Recite(untilGeneration int, lastSpoken int) (newLastSpoken int) {
	for m.generation < untilGeneration {
		lastSpoken = m.reciteOne(lastSpoken)
	}

	return lastSpoken
}

func (m *memory) reciteOne(lastSpoken int) (numberSpoken int) {
	defer m.remember(lastSpoken)

	if len(m.StartingNumbers) > 0 {
		return m.nextStartingNumber()
	}

	if iterationsAgo, ok := m.findPrevious(lastSpoken); ok {
		return iterationsAgo
	}

	// new number!
	return 0
}

func (m *memory) nextStartingNumber() (number int) {
	number = m.StartingNumbers[0]

	if len(m.StartingNumbers) > 1 {
		m.StartingNumbers = m.StartingNumbers[1:]
	} else {
		m.StartingNumbers = nil // no more starting numbers
	}

	return number
}

func (m *memory) findPrevious(number int) (generationsAgo int, ok bool) {
	// search through previous numbers, purposefully excluding the last entry
	if previousGeneration, ok := m.previousNumbers[number]; ok {
		generationsAgo = m.generation - previousGeneration
		return generationsAgo, true
	}

	return 0, false
}

func (m *memory) remember(newNumber int) {
	m.previousNumbers[newNumber] = m.generation
	m.generation++
}
