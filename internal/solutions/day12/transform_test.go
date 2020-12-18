package day12

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/segwin/adventofcode-2020/internal/geometry"
)

func TestTransformPart1(t *testing.T) {
	t.Parallel()

	type Test struct {
		// inputs
		Direction rune
		Magnitude int64
		Ship      geometry.Point
		Waypoint  geometry.Point

		// outputs
		ExpectedShip     geometry.Point
		ExpectedWaypoint geometry.Point
		ExpectedErr      error
	}

	testFn := func(t *testing.T, cfg Test) {
		t.Parallel()

		ship, waypoint, err := Transform(cfg.Direction, cfg.Magnitude, cfg.Ship, cfg.Waypoint, false)

		if got, expected := err, cfg.ExpectedErr; !errors.Is(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}

		if diff := cmp.Diff(cfg.ExpectedShip, ship); diff != "" {
			t.Errorf("Unexpected diff:\n%v", diff)
		}

		if diff := cmp.Diff(cfg.ExpectedWaypoint, waypoint); diff != "" {
			t.Errorf("Unexpected diff:\n%v", diff)
		}
	}

	tests := map[string]Test{
		"F10, prev=East": {
			Direction: 'F',
			Magnitude: 10,

			Ship:             geometry.NewInts(0, 0),
			Waypoint:         geometry.NewInts(1, 0),
			ExpectedShip:     geometry.NewInts(10, 0),
			ExpectedWaypoint: geometry.NewInts(1, 0),
		},

		"N3, prev=East": {
			Direction: 'N',
			Magnitude: 3,

			Ship:             geometry.NewInts(10, 0),
			Waypoint:         geometry.NewInts(1, 0),
			ExpectedShip:     geometry.NewInts(10, 3),
			ExpectedWaypoint: geometry.NewInts(1, 0),
		},

		"F7, prev=East": {
			Direction: 'F',
			Magnitude: 7,

			Ship:             geometry.NewInts(10, 3),
			Waypoint:         geometry.NewInts(1, 0),
			ExpectedShip:     geometry.NewInts(17, 3),
			ExpectedWaypoint: geometry.NewInts(1, 0),
		},

		"R90, prev=East": {
			Direction: 'R',
			Magnitude: 90,

			Ship:             geometry.NewInts(17, 3),
			Waypoint:         geometry.NewInts(1, 0),
			ExpectedShip:     geometry.NewInts(17, 3),
			ExpectedWaypoint: geometry.NewInts(0, -1),
		},

		"F11, prev=South": {
			Direction: 'F',
			Magnitude: 11,

			Ship:             geometry.NewInts(17, 3),
			Waypoint:         geometry.NewInts(0, -1),
			ExpectedShip:     geometry.NewInts(17, -8),
			ExpectedWaypoint: geometry.NewInts(0, -1),
		},

		"L90, prev=South": {
			Direction: 'L',
			Magnitude: 90,

			Ship:             geometry.NewInts(17, -8),
			Waypoint:         geometry.NewInts(0, -1),
			ExpectedShip:     geometry.NewInts(17, -8),
			ExpectedWaypoint: geometry.NewInts(1, 0),
		},

		"R270, prev=East": {
			Direction: 'R',
			Magnitude: 270,

			Ship:             geometry.NewInts(17, -8),
			Waypoint:         geometry.NewInts(1, 0),
			ExpectedShip:     geometry.NewInts(17, -8),
			ExpectedWaypoint: geometry.NewInts(0, 1),
		},

		"L270, prev=North": {
			Direction: 'L',
			Magnitude: 270,

			Ship:             geometry.NewInts(17, -8),
			Waypoint:         geometry.NewInts(0, 1),
			ExpectedShip:     geometry.NewInts(17, -8),
			ExpectedWaypoint: geometry.NewInts(1, 0),
		},
	}

	for name, cfg := range tests {
		t.Run(name, func(t *testing.T) { testFn(t, cfg) })
	}
}

func TestTransformPart2(t *testing.T) {
	t.Parallel()

	type Test struct {
		// inputs
		Direction rune
		Magnitude int64
		Ship      geometry.Point
		Waypoint  geometry.Point

		// outputs
		ExpectedShip     geometry.Point
		ExpectedWaypoint geometry.Point
		ExpectedErr      error
	}

	testFn := func(t *testing.T, cfg Test) {
		t.Parallel()

		ship, waypoint, err := Transform(cfg.Direction, cfg.Magnitude, cfg.Ship, cfg.Waypoint, true)

		if got, expected := err, cfg.ExpectedErr; !errors.Is(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}

		if diff := cmp.Diff(cfg.ExpectedShip, ship); diff != "" {
			t.Errorf("Unexpected diff:\n%v", diff)
		}

		if diff := cmp.Diff(cfg.ExpectedWaypoint, waypoint); diff != "" {
			t.Errorf("Unexpected diff:\n%v", diff)
		}
	}

	tests := map[string]Test{
		"F10, prev=East": {
			Direction: 'F',
			Magnitude: 10,

			Ship:             geometry.NewInts(0, 0),
			Waypoint:         geometry.NewInts(10, 1),
			ExpectedShip:     geometry.NewInts(100, 10),
			ExpectedWaypoint: geometry.NewInts(10, 1),
		},

		"N3, prev=East": {
			Direction: 'N',
			Magnitude: 3,

			Ship:             geometry.NewInts(100, 10),
			Waypoint:         geometry.NewInts(10, 1),
			ExpectedShip:     geometry.NewInts(100, 10),
			ExpectedWaypoint: geometry.NewInts(10, 4),
		},

		"F7, prev=East": {
			Direction: 'F',
			Magnitude: 7,

			Ship:             geometry.NewInts(100, 10),
			Waypoint:         geometry.NewInts(10, 4),
			ExpectedShip:     geometry.NewInts(170, 38),
			ExpectedWaypoint: geometry.NewInts(10, 4),
		},

		"R90, prev=East": {
			Direction: 'R',
			Magnitude: 90,

			Ship:             geometry.NewInts(170, 38),
			Waypoint:         geometry.NewInts(10, 4),
			ExpectedShip:     geometry.NewInts(170, 38),
			ExpectedWaypoint: geometry.NewInts(4, -10),
		},

		"F11, prev=South": {
			Direction: 'F',
			Magnitude: 11,

			Ship:             geometry.NewInts(170, 38),
			Waypoint:         geometry.NewInts(4, -10),
			ExpectedShip:     geometry.NewInts(214, -72),
			ExpectedWaypoint: geometry.NewInts(4, -10),
		},
	}

	for name, cfg := range tests {
		t.Run(name, func(t *testing.T) { testFn(t, cfg) })
	}
}
