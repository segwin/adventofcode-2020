package day14

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrMismatchedBitsetLengths = errors.New("mismatched bitset lengths")
	ErrBitsetOverflow          = errors.New("bitset operation caused an overflow")
)

type Bit byte

const (
	Zero Bit = 0
	One  Bit = 1
)

func (b Bit) String() string {
	switch b {
	case Zero:
		return "0"
	case One:
		return "1"
	default:
		return fmt.Sprintf("<invalid: %d>", b)
	}
}

// Bitset stores a sequence of bits and allows conversion to/from integer values.
// The internal representation is little endian.
type Bitset struct {
	bits []Bit
}

func NewBitset(bitSize int) *Bitset {
	return &Bitset{bits: make([]Bit, bitSize)}
}

func (b *Bitset) Clone() *Bitset {
	clone := &Bitset{bits: make([]Bit, b.Len())}
	copy(clone.bits, b.bits)
	return clone
}

func (b *Bitset) Unmarshal(encodedInteger string, bitSize int) error {
	value, err := strconv.ParseInt(encodedInteger, 10, bitSize)
	if err != nil {
		return err
	}

	b.ParseInt(value, 36) // size assumption is always valid for part 1
	return nil
}

func (b *Bitset) Len() int { return len(b.bits) }

// ParseInt creates a binary representation of this integer and stores it in this
// bitset.
func (b *Bitset) ParseInt(value int64, bitSize int) {
	b.bits = make([]Bit, bitSize)
	for i := range b.bits {
		b.bits[i] = Bit((value >> i) & 0b1)
	}
}

// // Add the given bitset to this one. Both bitsets must have the same length or an
// // error is returned.
// func (b *Bitset) Add(other *Bitset) error {
// 	if other.Len() > b.Len() {
// 		return ErrMismatchedBitsetLengths
// 	}

// 	carryOver := Zero
// 	for i, added := range other.bits {
// 		b.bits[i] += added + carryOver // add new value & previous carry-over
// 		carryOver = b.bits[i] >> 1     // take >LSB as carry-over value
// 		b.bits[i] &= 0b1               // only keep LSB in this bit value
// 	}

// 	return nil
// }

// Int returns the integer representation of this bitset. It overflows if the bitset
// is larger than 64 bits.
func (b *Bitset) Int() (value int64) {
	for i := 0; i < b.Len(); i++ {
		value += int64(b.bits[i]) << i
	}

	return value
}

// String returns the string representation of this bitset, i.e. the sequence of bits.
func (b *Bitset) String() (value string) {
	// navigate in reverse order as bitset itself is little-endian
	for i := b.Len() - 1; i >= 0; i-- {
		value += b.bits[i].String()
	}

	return value
}
