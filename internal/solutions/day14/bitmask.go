package day14

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidBitmaskCharacter = errors.New("invalid character in encoded bitmask")
	ErrBitmaskTooLong          = errors.New("bitmask is more than 36 bits long")
	ErrInvalidBitPermutations  = errors.New("got invalid number of bit permutations")
)

type BitSetter func(currentBit Bit) (newBits []Bit)

var (
	NoChange BitSetter = func(currentBit Bit) []Bit { return []Bit{currentBit} }

	// part 1 setters
	AssignZero BitSetter = func(_ Bit) []Bit { return []Bit{Zero} }
	AssignOne  BitSetter = func(_ Bit) []Bit { return []Bit{One} }

	// part 2 setters
	AssignWildcard BitSetter = func(currentBit Bit) []Bit { return []Bit{Zero, One} }
)

type Bitmask struct {
	bitSetters []BitSetter
}

func (b *Bitmask) Unmarshal(encodedMask string, part1 bool) error {
	b.bitSetters = make([]BitSetter, 0, 36) // size assumption is always valid for part 1
	if len(encodedMask) > cap(b.bitSetters) {
		return fmt.Errorf("%w (%d)", ErrBitmaskTooLong, len(encodedMask))
	}

	// reverse iterate through encoded mask: it's encoded left-to-right but bitsets
	// are little-endian (right-to-left)
	for i := len(encodedMask) - 1; i >= 0; i-- {
		setter := NoChange

		switch value := encodedMask[i]; value {
		case '0':
			if part1 {
				setter = AssignZero
			}

		case '1':
			setter = AssignOne

		case 'X':
			if !part1 {
				setter = AssignWildcard
			}

		default:
			return fmt.Errorf("%w (%s)", ErrInvalidBitmaskCharacter, string(value))
		}

		b.bitSetters = append(b.bitSetters, setter)
	}

	return nil
}

func (b *Bitmask) Len() int { return len(b.bitSetters) }

func (b *Bitmask) Apply(bits *Bitset) (maskedBitsets []*Bitset, err error) {
	if bits.Len() != b.Len() {
		return nil, ErrMismatchedBitsetLengths
	}

	maskedBitsets = []*Bitset{
		// start with only the original bitset, we'll add more permutations as they come up
		bits.Clone(),
	}

	for bitIdx, mutate := range b.bitSetters {
		for _, bitset := range maskedBitsets {
			newBits := mutate(bitset.bits[bitIdx])

			if len(newBits) < 1 || len(newBits) > 2 {
				// bit permutations can only affect the current bitset and optionally add 1 new permutation
				return nil, ErrInvalidBitPermutations
			}

			bitset.bits[bitIdx] = newBits[0]
			if len(newBits) == 2 {
				newBitset := bitset.Clone()
				newBitset.bits[bitIdx] = newBits[1]

				maskedBitsets = append(maskedBitsets, newBitset)
			}
		}
	}

	return maskedBitsets, nil
}
