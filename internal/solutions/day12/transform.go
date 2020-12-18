package day12

import (
	"errors"
	"fmt"

	"github.com/segwin/adventofcode-2020/internal/geometry"
)

var (
	ErrInvalidDirection = errors.New("invalid direction")
)

func Transform(direction rune, magnitude int64, ship, waypoint geometry.Point, part2 bool) (newShip, newWaypoint geometry.Point, err error) {
	cardinalDirection, ok := ToCardinalDirection(direction)
	if ok {
		newShip, newWaypoint = cardinalTransform(cardinalDirection, magnitude, ship, waypoint, part2)
		return newShip, newWaypoint, nil
	}

	relativeDirection, ok := ToRelativeDirection(direction)
	if ok {
		newShip, newWaypoint = relativeTransform(relativeDirection, magnitude, ship, waypoint)
		return newShip, newWaypoint, nil
	}

	return nil, nil, fmt.Errorf("%w (%s%d)", ErrInvalidDirection, string(direction), magnitude)
}

func cardinalTransform(direction CardinalDirection, magnitude int64, ship, waypoint geometry.Point, part2 bool) (newShip, newWaypoint geometry.Point) {
	if part2 {
		// part 2: translate waypoint
		return ship, direction.Translation(magnitude).Apply(waypoint)
	}

	// part 1: translate ship
	return direction.Translation(magnitude).Apply(ship), waypoint
}

func relativeTransform(direction RelativeDirection, magnitude int64, ship, waypoint geometry.Point) (newShip, newWaypoint geometry.Point) {
	transform := direction.Transformation(magnitude, waypoint)

	// we never receive both a ship & waypoint transform in this problem, so let's
	// not worry about operation order here
	transformsShip, transformsWaypoint := direction.Transforms()
	if transformsShip {
		ship = transform.Apply(ship)
	}
	if transformsWaypoint {
		waypoint = transform.Apply(waypoint)
	}

	return ship, waypoint
}

type Transformation interface {
	Apply(point geometry.Point) (newPoint geometry.Point)
}

type Translation struct {
	coord geometry.Point
}

// Apply moves a point by the amounts given in this Translation.
func (t *Translation) Apply(point geometry.Point) (newPoint geometry.Point) {
	if t == nil {
		return point
	}

	return geometry.NewInts(
		point.MustGet(0).Int()+t.coord.MustGet(0).Int(),
		point.MustGet(1).Int()+t.coord.MustGet(1).Int(),
	)
}

type Rotation struct {
	Degrees int64
}

// Apply rotates a point around the global origin (0,0). Only right angle rotations
// are supported (i.e. rotations that are multiples of 90 degrees).
func (r *Rotation) Apply(point geometry.Point) (newPoint geometry.Point) {
	if r == nil {
		return point
	}

	if r.Degrees%90 != 0 {
		// no need to handle non-90 degree rotations for this problem, let's not
		// bother with error handling here
		panic("rotations that are not multiples of 90 are not supported")
	}

	switch (r.Degrees + 360) % 360 {
	default: // 0
		return point
	case 90:
		return geometry.NewInts(point.MustGet(1).Int(), -point.MustGet(0).Int())
	case 180:
		return geometry.NewInts(-point.MustGet(0).Int(), -point.MustGet(1).Int())
	case 270:
		return geometry.NewInts(-point.MustGet(1).Int(), point.MustGet(0).Int())
	}
}
