package day12

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTransformPart1(t *testing.T) {
	t.Parallel()

	type Test struct {
		// inputs
		Direction rune
		Magnitude int
		Ship      *Coordinates
		Waypoint  *Coordinates

		// outputs
		ExpectedShip     *Coordinates
		ExpectedWaypoint *Coordinates
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

			Ship:             &Coordinates{},
			Waypoint:         &Coordinates{X: 1},
			ExpectedShip:     &Coordinates{X: 10},
			ExpectedWaypoint: &Coordinates{X: 1},
		},

		"N3, prev=East": {
			Direction: 'N',
			Magnitude: 3,

			Ship:             &Coordinates{X: 10},
			Waypoint:         &Coordinates{X: 1},
			ExpectedShip:     &Coordinates{X: 10, Y: 3},
			ExpectedWaypoint: &Coordinates{X: 1},
		},

		"F7, prev=East": {
			Direction: 'F',
			Magnitude: 7,

			Ship:             &Coordinates{X: 10, Y: 3},
			Waypoint:         &Coordinates{X: 1},
			ExpectedShip:     &Coordinates{X: 17, Y: 3},
			ExpectedWaypoint: &Coordinates{X: 1},
		},

		"R90, prev=East": {
			Direction: 'R',
			Magnitude: 90,

			Ship:             &Coordinates{X: 17, Y: 3},
			Waypoint:         &Coordinates{X: 1},
			ExpectedShip:     &Coordinates{X: 17, Y: 3},
			ExpectedWaypoint: &Coordinates{Y: -1},
		},

		"F11, prev=South": {
			Direction: 'F',
			Magnitude: 11,

			Ship:             &Coordinates{X: 17, Y: 3},
			Waypoint:         &Coordinates{Y: -1},
			ExpectedShip:     &Coordinates{X: 17, Y: -8},
			ExpectedWaypoint: &Coordinates{Y: -1},
		},

		"L90, prev=South": {
			Direction: 'L',
			Magnitude: 90,

			Ship:             &Coordinates{X: 17, Y: -8},
			Waypoint:         &Coordinates{Y: -1},
			ExpectedShip:     &Coordinates{X: 17, Y: -8},
			ExpectedWaypoint: &Coordinates{X: 1},
		},

		"R270, prev=East": {
			Direction: 'R',
			Magnitude: 270,

			Ship:             &Coordinates{X: 17, Y: -8},
			Waypoint:         &Coordinates{X: 1},
			ExpectedShip:     &Coordinates{X: 17, Y: -8},
			ExpectedWaypoint: &Coordinates{Y: 1},
		},

		"L270, prev=North": {
			Direction: 'L',
			Magnitude: 270,

			Ship:             &Coordinates{X: 17, Y: -8},
			Waypoint:         &Coordinates{Y: 1},
			ExpectedShip:     &Coordinates{X: 17, Y: -8},
			ExpectedWaypoint: &Coordinates{X: 1},
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
		Magnitude int
		Ship      *Coordinates
		Waypoint  *Coordinates

		// outputs
		ExpectedShip     *Coordinates
		ExpectedWaypoint *Coordinates
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

			Ship:             &Coordinates{},
			Waypoint:         &Coordinates{X: 10, Y: 1},
			ExpectedShip:     &Coordinates{X: 100, Y: 10},
			ExpectedWaypoint: &Coordinates{X: 10, Y: 1},
		},

		"N3, prev=East": {
			Direction: 'N',
			Magnitude: 3,

			Ship:             &Coordinates{X: 100, Y: 10},
			Waypoint:         &Coordinates{X: 10, Y: 1},
			ExpectedShip:     &Coordinates{X: 100, Y: 10},
			ExpectedWaypoint: &Coordinates{X: 10, Y: 4},
		},

		"F7, prev=East": {
			Direction: 'F',
			Magnitude: 7,

			Ship:             &Coordinates{X: 100, Y: 10},
			Waypoint:         &Coordinates{X: 10, Y: 4},
			ExpectedShip:     &Coordinates{X: 170, Y: 38},
			ExpectedWaypoint: &Coordinates{X: 10, Y: 4},
		},

		"R90, prev=East": {
			Direction: 'R',
			Magnitude: 90,

			Ship:             &Coordinates{X: 170, Y: 38},
			Waypoint:         &Coordinates{X: 10, Y: 4},
			ExpectedShip:     &Coordinates{X: 170, Y: 38},
			ExpectedWaypoint: &Coordinates{X: 4, Y: -10},
		},

		"F11, prev=South": {
			Direction: 'F',
			Magnitude: 11,

			Ship:             &Coordinates{X: 170, Y: 38},
			Waypoint:         &Coordinates{X: 4, Y: -10},
			ExpectedShip:     &Coordinates{X: 214, Y: -72},
			ExpectedWaypoint: &Coordinates{X: 4, Y: -10},
		},
	}

	for name, cfg := range tests {
		t.Run(name, func(t *testing.T) { testFn(t, cfg) })
	}
}
