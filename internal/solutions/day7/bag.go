package day7

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidMsg   = errors.New("invalid bag message")
	ErrInvalidCount = errors.New("invalid count in allowed bags list")
)

type Bag interface {
	// Unmarshal the given message into this Bag object.
	Unmarshal(msg string) error

	// Colour returns the colour of this bag.
	Colour() string

	// Contains returns true if this bag or any of its sub-bags contain a bag of
	// the given colour.
	Contains(colour string, allBags map[string]Bag) bool

	// SubBags returns all sub-bags contained within this bag, mapped to their
	// count.
	SubBags(allBags map[string]Bag) map[Bag]int64
}

type bag struct {
	colour   string
	contains map[string]int64 // map of colour -> allowed count
}

func (b *bag) Unmarshal(msg string) error {
	msgHalves := strings.SplitN(msg, " contain ", 2)
	if len(msgHalves) != 2 {
		return fmt.Errorf("%w: expected \"a contain x, y, z\"", ErrInvalidMsg)
	}

	b.colour = strings.TrimSuffix(strings.TrimSuffix(msgHalves[0], "s"), " bag")
	b.contains = map[string]int64{}

	allowedColoursStr := strings.TrimSuffix(strings.TrimSpace(msgHalves[1]), ".")
	if allowedColoursStr == "no other bags" {
		return nil // we're done
	}

	allowedColourMsgs := strings.Split(allowedColoursStr, ", ")
	for _, allowedColourMsg := range allowedColourMsgs {
		countAndColour := strings.SplitN(allowedColourMsg, " ", 2)
		if len(countAndColour) != 2 {
			return fmt.Errorf("%w: expected \"a contain x, y, z\"", ErrInvalidMsg)
		}

		count, err := strconv.Atoi(countAndColour[0])
		if err != nil {
			return fmt.Errorf("%w: bad count in %q", ErrInvalidCount, allowedColourMsg)
		}

		colour := strings.TrimSuffix(strings.TrimSuffix(countAndColour[1], "s"), " bag")
		b.contains[colour] = int64(count)
	}

	return nil
}

func (b *bag) Colour() string {
	return b.colour
}

func (b *bag) Contains(colour string, allBags map[string]Bag) bool {
	if count, ok := b.contains[colour]; ok && count > 0 {
		return true
	}

	for subBagColour, count := range b.contains {
		if count == 0 {
			continue
		}

		subBag, ok := allBags[subBagColour]
		if !ok {
			continue // no such bag is known
		}

		if subBag.Contains(colour, allBags) {
			return true
		}
	}

	return false
}

func (b *bag) SubBags(allBags map[string]Bag) (subBags map[Bag]int64) {
	subBags = make(map[Bag]int64)

	for subBagColour, subBagCount := range b.contains {
		subBag, ok := allBags[subBagColour]
		if !ok {
			continue
		}

		subBags[subBag] += subBagCount

		subSubBags := subBag.SubBags(allBags)
		for subSubBag, subSubBagCount := range subSubBags {
			subBags[subSubBag] += subBagCount * subSubBagCount
		}
	}

	return subBags
}
