package day16

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInvalidTicketFieldLine = errors.New("invalid ticket field line")
	ErrInvalidTicketLine      = errors.New("invalid ticket line")
)

type TicketField struct {
	Name   string
	Ranges []Range
}

func (f *TicketField) Unmarshal(line string) error {
	matches := regexp.
		MustCompile("([^:]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)").
		FindStringSubmatch(line)

	if len(matches) != 1+5 {
		return fmt.Errorf("%w (%q)", ErrInvalidTicketFieldLine, line)
	}

	// get name
	name := matches[1]

	// build ranges
	var rangeValues []int
	for i := 2; i < len(matches); i++ {
		rangeValue, err := strconv.Atoi(matches[i])
		if err != nil {
			return err
		}

		rangeValues = append(rangeValues, rangeValue)
	}

	ranges := make([]Range, 0, len(rangeValues)/2)
	for i := 0; i < len(rangeValues); i += 2 {
		ranges = append(ranges, Range{Min: rangeValues[i], Max: rangeValues[i+1]})
	}

	// populate ticket field now that all components have been validated
	f.Name = name
	f.Ranges = ranges

	return nil
}

type RawTicket struct {
	Values []int
}

func (t *RawTicket) Unmarshal(line string) error {
	strValues := strings.Split(line, ",")

	values := make([]int, len(strValues))
	for i, strValue := range strValues {
		value, err := strconv.Atoi(strValue)
		if err != nil {
			return err
		}

		values[i] = value
	}

	// populate ticket now that values have been validated
	t.Values = values

	return nil
}

func (t *RawTicket) InvalidValues(fields []*TicketField) (invalidValues []int) {
	takenBy := map[*TicketField]int{}
	for i, value := range t.Values {
		isValid := false

		for _, field := range fields {
			if _, ok := takenBy[field]; ok {
				// field already used
				continue
			}

			for _, fieldRange := range field.Ranges {
				if fieldRange.Contains(value) {
					takenBy[field] = i
					isValid = true
					break
				}
			}

			if isValid {
				break
			}
		}

		if !isValid {
			invalidValues = append(invalidValues, value)
		}
	}

	return invalidValues
}

type PotentialPositions map[*TicketField][]int

func NewPotentialPositions(fields []*TicketField, numPositions int) PotentialPositions {
	potentialPositions := make(PotentialPositions)
	for _, field := range fields {
		for i := 0; i < numPositions; i++ {
			potentialPositions[field] = append(potentialPositions[field], i)
		}
	}

	return potentialPositions
}

func (p PotentialPositions) Intersect(fields []*TicketField, ticket *RawTicket) {
	for field, positions := range p {
		var validPositions []int
		for _, position := range positions {
			for _, fieldRange := range field.Ranges {
				if fieldRange.Contains(ticket.Values[position]) {
					validPositions = append(validPositions, position)
					break
				}
			}
		}

		// prune all invalid positions for this field
		p[field] = validPositions
	}
}

// Collapse tries to collapse all potential positions in this list into a single
// possibility per field. For this algorithm to work, the following conditions must
// be met:
//
//   - Each field must map to a unique number of potential positions
//   - There must be exactly one field with only 1 possible position and one with
//     all possible positions (e.g. fieldN=0..N, field0=<single value>)
//   - A field with N positions must contain all the same possibilities as the one
//     with N-1 positions, plus one new possibility (e.g. field3=>0,3,5, field2=>0,5,
//     field1=>5)
func (p PotentialPositions) Collapse() map[*TicketField]int {
	collapsedPositions := map[*TicketField]int{}

	var current, next *TicketField
	for i := 0; i < len(p); i++ {
		if current == nil {
			for field, positions := range p {
				if len(positions) == 1 {
					// start collapse here
					current = field
				}
			}
		}

		collapsedPositions[current] = p[current][0]
		for field := range p {
			// collapse: remove current position from list of possibilities
			p[field] = filterOut(p[field], collapsedPositions[current])

			if len(p[field]) == 1 {
				next = field // found next field to collapse
			}
		}

		current = next
		next = nil
	}

	return collapsedPositions
}

func filterOut(values []int, remove int) (newValues []int) {
	for _, value := range values {
		if value != remove {
			newValues = append(newValues, value)
		}
	}

	return newValues
}
