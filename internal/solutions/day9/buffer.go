package day9

import (
	"errors"
	"fmt"
)

type Ring interface {
	// Len returns the length of this ring buffer, i.e. the number of elements
	// it currently contains.
	Len() int

	// Cap returns the capacity of this ring buffer, i.e. the number of elements
	// it *can* contain.
	Cap() int

	// Push sets the value at the current position in the ring buffer and advances
	// to the next position. If a value already exists at this position, it is
	// overwritten.
	Push(value interface{}) error

	// Get returns the value at position i in the ring buffer. An error is returned
	// if the operation fails.
	Get(i int) (value interface{}, err error)

	// MustGet returns the value at position i in the ring buffer. If any error
	// occurs, the program will panic.
	MustGet(i int) (value interface{})
}

var (
	ErrInvalidRingValue = errors.New("invalid value for this ring")
	ErrOutOfRange       = errors.New("index out of range")
)

type intRing struct {
	values   []int
	position int
}

func NewIntRing(capacity int) Ring {
	return &intRing{
		values: make([]int, 0, capacity),
	}
}

func (r *intRing) Len() int { return len(r.values) }
func (r *intRing) Cap() int { return cap(r.values) }

func (r *intRing) Push(v interface{}) error {
	value, ok := v.(int)
	if !ok {
		return fmt.Errorf("%w (expects int, got %T)", ErrInvalidRingValue, v)
	}

	if len(r.values) <= r.position {
		// first run around the ring, this position hasn't yet been initialised
		r.values = append(r.values, value)
	} else {
		r.values[r.position] = value
	}

	r.next()

	return nil
}

func (r *intRing) next() {
	r.position++
	r.position %= r.Cap()
}

func (r *intRing) Get(i int) (value interface{}, err error) {
	if length := r.Len(); i >= length {
		return nil, fmt.Errorf("%w (%d requested for buffer of size %d)", ErrOutOfRange, i, length)
	}

	return r.values[i], nil
}

func (r *intRing) MustGet(i int) (value interface{}) {
	value, err := r.Get(i)
	if err != nil {
		panic(err)
	}

	return value
}
