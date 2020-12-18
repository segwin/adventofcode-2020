package geometry

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidDimension = errors.New("invalid dimension value")
)

type Point interface {
	// Get returns the current value for the given dimension. An error is returned if
	// the dimension is greater than the number of dimensions for this coordinate, an
	// error is returned.
	Get(dimension int) (Number, error)

	// MustSet returns the current value for the given dimension. It panics if an error
	// occurs (e.g. dimension greater than number of dimensions in coordinates).
	MustGet(dimension int) Number

	// Set assigns a value to the given dimension. An error is returned if the dimension
	// is greater than the number of dimensions for this coordinate, an error is
	// returned.
	Set(value Number, dimension int) error

	// MustSet assigns a value to the given dimension. It panics if an error occurs
	// (e.g. dimension greater than number of dimensions in coordinates).
	MustSet(value Number, dimension int)
}

func newPoint(numDimensions int) Point {
	return &coordinates{make([]Number, numDimensions)}
}

func NewInts(values ...int64) Point {
	coordinates := newPoint(len(values))
	for i, value := range values {
		coordinates.MustSet(Int(value), i)
	}

	return coordinates
}

func NewFloats(values ...float64) Point {
	coordinates := newPoint(len(values))
	for i, value := range values {
		coordinates.MustSet(Float(value), i)
	}

	return coordinates
}

type coordinates struct {
	Dimensions []Number
}

func (c *coordinates) Get(dimension int) (Number, error) {
	if err := c.validate(dimension); err != nil {
		return nil, err
	}

	return c.Dimensions[dimension], nil
}

func (c *coordinates) MustGet(dimension int) Number {
	value, err := c.Get(dimension)
	if err != nil {
		panic(err)
	}

	return value
}

func (c *coordinates) Set(value Number, dimension int) error {
	if err := c.validate(dimension); err != nil {
		return err
	}

	c.Dimensions[dimension] = value
	return nil
}

func (c *coordinates) MustSet(value Number, dimension int) {
	if err := c.Set(value, dimension); err != nil {
		panic(err)
	}
}

func (c *coordinates) validate(dimension int) error {
	if dimension > len(c.Dimensions) {
		return fmt.Errorf("%w (got %d for %dD coordinates", ErrInvalidDimension, dimension, len(c.Dimensions))
	}

	return nil
}
