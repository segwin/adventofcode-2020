package day12

import "math"

type CardinalDirection uint

const (
	// cardinal directions
	North CardinalDirection = 0
	East  CardinalDirection = 1
	South CardinalDirection = 2
	West  CardinalDirection = 3

	numCardinalDirections = 4
)

func (d CardinalDirection) Translation(magnitude int) *Translation {
	switch d {
	case North:
		return &Translation{Y: +magnitude}
	case South:
		return &Translation{Y: -magnitude}
	case East:
		return &Translation{X: +magnitude}
	case West:
		return &Translation{X: -magnitude}
	}

	return nil
}

func (d CardinalDirection) RotationFrom(reference CardinalDirection) *Rotation {
	// wrap difference into 0-4 range
	numRotations := uint(reference-d) % numCardinalDirections
	if numRotations == 0 {
		return nil
	}

	return &Rotation{Degrees: int(numRotations) * 90}
}

func (d CardinalDirection) String() string {
	switch d {
	case North:
		return "N"
	case South:
		return "S"
	case East:
		return "E"
	case West:
		return "W"
	}

	return "<unknown>"
}

func ToCardinalDirection(direction rune) (cardinalDirection CardinalDirection, ok bool) {
	switch direction {
	case 'n', 'N':
		return North, true
	case 's', 'S':
		return South, true
	case 'e', 'E':
		return East, true
	case 'w', 'W':
		return West, true
	}

	return math.MaxUint32, false
}

type RelativeDirection rune

const (
	// turn directions
	Left  RelativeDirection = -1
	Right RelativeDirection = +1

	// advance directions
	Forward RelativeDirection = 0
)

func (d RelativeDirection) Transforms() (ship, waypoint bool) {
	switch d {
	case Left, Right:
		return false, true
	case Forward:
		return true, false
	}

	return false, false
}

func (d RelativeDirection) Transformation(magnitude int, waypoint *Coordinates) Transformation {
	switch d {
	case Left, Right:
		return &Rotation{
			Degrees: int(d) * magnitude,
		}

	case Forward:
		return &Translation{
			X: waypoint.X * magnitude,
			Y: waypoint.Y * magnitude,
		}
	}

	return nil
}

func ToRelativeDirection(direction rune) (relativeDirection RelativeDirection, ok bool) {
	switch direction {
	case 'L':
		return Left, true
	case 'R':
		return Right, true
	case 'F':
		return Forward, true
	}

	return -1, false
}
