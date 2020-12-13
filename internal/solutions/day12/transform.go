package day12

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidDirection = errors.New("invalid direction")
)

func Transform(direction rune, magnitude int, ship, waypoint *Coordinates, part2 bool) (newShip, newWaypoint *Coordinates, err error) {
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

func cardinalTransform(direction CardinalDirection, magnitude int, ship, waypoint *Coordinates, part2 bool) (newShip, newWaypoint *Coordinates) {
	if part2 {
		// part 2: translate waypoint
		return ship, direction.Translation(magnitude).Apply(waypoint)
	}

	// part 1: translate ship
	return direction.Translation(magnitude).Apply(ship), waypoint
}

func relativeTransform(direction RelativeDirection, magnitude int, ship, waypoint *Coordinates) (newShip, newWaypoint *Coordinates) {
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

type Coordinates struct {
	X int
	Y int
}

type Transformation interface {
	Apply(point *Coordinates) (newPoint *Coordinates)
}

type Translation struct {
	X int
	Y int
}

// Apply moves a point by the amounts given in this Translation.
func (t *Translation) Apply(point *Coordinates) (newPoint *Coordinates) {
	if t == nil {
		return point
	}

	return &Coordinates{
		X: point.X + t.X,
		Y: point.Y + t.Y,
	}
}

type Rotation struct {
	Degrees int
}

// Apply rotates a point around the global origin (0,0). Only right angle rotations
// are supported (i.e. rotations that are multiples of 90 degrees).
func (r *Rotation) Apply(point *Coordinates) (newPoint *Coordinates) {
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
		return &Coordinates{X: point.Y, Y: -point.X}
	case 180:
		return &Coordinates{X: -point.X, Y: -point.Y}
	case 270:
		return &Coordinates{X: -point.Y, Y: point.X}
	}
}
