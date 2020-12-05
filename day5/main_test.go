package main

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnmarshalSeat(t *testing.T) {
	t.Parallel()

	type Test struct {
		// inputs
		encoded string

		// outputs
		expectedSeat Seat
		expectedID   uint16
		expectedErr  error
	}

	testFn := func(t *testing.T, cfg Test) {
		t.Parallel()

		seat := Seat{}
		err := seat.Unmarshal(cfg.encoded)
		if !errors.Is(err, cfg.expectedErr) {
			t.Errorf("Got %v, expected %v", err, cfg.expectedErr)
		}

		if err != nil {
			return // we're done
		}

		if diff := cmp.Diff(cfg.expectedSeat, seat); diff != "" {
			t.Errorf("Unexpected diff:\n%v", diff)
		}

		if got, expected := seat.ID(), cfg.expectedID; got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	}

	tests := map[string]Test{
		"ok: FBFBBFFRLR": {
			encoded:      "FBFBBFFRLR",
			expectedSeat: Seat{Row: 44, Column: 5},
			expectedID:   357,
		},

		"ok: BFFFBBFRRR": {
			encoded:      "BFFFBBFRRR",
			expectedSeat: Seat{Row: 70, Column: 7},
			expectedID:   567,
		},

		"ok: FFFBBBFRRR": {
			encoded:      "FFFBBBFRRR",
			expectedSeat: Seat{Row: 14, Column: 7},
			expectedID:   119,
		},

		"ok: BBFFBBFRLL": {
			encoded:      "BBFFBBFRLL",
			expectedSeat: Seat{Row: 102, Column: 4},
			expectedID:   820,
		},
	}

	for name, cfg := range tests {
		t.Run(name, func(t *testing.T) { testFn(t, cfg) })
	}
}
