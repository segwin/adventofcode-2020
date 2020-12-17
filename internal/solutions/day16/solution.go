package day16

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/segwin/adventofcode-2020/internal/input"
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to create input file scanner: %v\n", err)
		os.Exit(1)
	}

	fields, myTicket, otherTickets, err := s.parseLines(scanner)
	if err != nil {
		fmt.Printf("ERROR: Failed to get problem values: %v\n", err)
		os.Exit(1)
	}

	validTickets := s.part1(fields, otherTickets)
	s.part2(fields, myTicket, validTickets)
}

func (s *Solution) parseLines(scanner input.Scanner) (fields []*TicketField, myTicket *RawTicket, otherTickets []*RawTicket, err error) {
	stage := Rules

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, nil, nil, err
		}

		line := strings.TrimSpace(scanner.Text())

		switch line {
		case "your ticket:":
			stage = MyTicket
			continue

		case "nearby tickets:":
			stage = OtherTickets
			continue

		case "":
			continue
		}

		switch stage {
		case Rules:
			field := &TicketField{}
			if err := field.Unmarshal(line); err != nil {
				return nil, nil, nil, err
			}

			fields = append(fields, field)

		case MyTicket:
			myTicket = &RawTicket{}
			if err := myTicket.Unmarshal(line); err != nil {
				return nil, nil, nil, err
			}

		case OtherTickets:
			ticket := &RawTicket{}
			if err := ticket.Unmarshal(line); err != nil {
				return nil, nil, nil, err
			}

			otherTickets = append(otherTickets, ticket)
		}
	}

	return fields, myTicket, otherTickets, nil
}

func (s *Solution) part1(fields []*TicketField, tickets []*RawTicket) (validTickets []*RawTicket) {
	fmt.Println("\nPART 1")

	errorRate := 0
	for _, ticket := range tickets {
		invalidValues := ticket.InvalidValues(fields)
		if len(invalidValues) == 0 {
			validTickets = append(validTickets, ticket)
		}

		for _, invalidValue := range invalidValues {
			errorRate += invalidValue
		}
	}

	fmt.Printf("  RESULT: Got an error rate of %d\n", errorRate)
	return validTickets
}

func (s *Solution) part2(fields []*TicketField, myTicket *RawTicket, validTickets []*RawTicket) {
	fmt.Println("\nPART 2")

	// initialise map of potential positions with all possibilities
	numPositions := len(myTicket.Values)
	potentialPositions := NewPotentialPositions(fields, numPositions)

	// narrow down possibilities, ticket by ticket
	for _, ticket := range validTickets {
		potentialPositions.Intersect(fields, ticket)
	}

	potentialPositions.Intersect(fields, myTicket)

	fmt.Println()
	for field, positions := range potentialPositions {
		positionStrings := make([]string, len(positions))
		for i, position := range positions {
			positionStrings[i] = strconv.Itoa(position)
		}

		fmt.Printf("%s => %s\n", field.Name, strings.Join(positionStrings, ", "))
	}
	fmt.Println()

	fieldPositions := potentialPositions.Collapse()

	product := 1
	for field, position := range fieldPositions {
		if !strings.HasPrefix(field.Name, "departure") {
			continue
		}

		product *= myTicket.Values[position]
	}

	fmt.Printf("  RESULT: The product of all departure fields on my ticket is %d\n", product)
}
