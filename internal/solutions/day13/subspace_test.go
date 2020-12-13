package day13

import (
	"errors"
	"testing"
)

func TestSubspaceIntersect(t *testing.T) {
	type Test struct {
		A *Subspace
		B *Subspace

		Expected    *Subspace
		ExpectedErr error
	}

	testFn := func(t *testing.T, cfg Test) {
		intersection, err := cfg.A.Intersect(cfg.B)
		if got, expected := err, cfg.ExpectedErr; !errors.Is(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}

		if err != nil || cfg.ExpectedErr != nil {
			return // we're done
		}

		if got, expected := intersection.Coefficient.Int64(), cfg.Expected.Coefficient.Int64(); got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}

		if got, expected := intersection.Offset.Int64(), cfg.Expected.Offset.Int64(); got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	}

	tests := map[string]Test{
		"2N ∩ 2N => 2N": {
			A: NewSubspace(2, 0),
			B: NewSubspace(2, 0),

			Expected: NewSubspace(2, 0),
		},

		"1N ∩ 2N => 2N": {
			A: NewSubspace(1, 0),
			B: NewSubspace(2, 0),

			Expected: NewSubspace(2, 0),
		},

		"2N ∩ 1N => 2N": {
			A: NewSubspace(2, 0),
			B: NewSubspace(1, 0),

			Expected: NewSubspace(2, 0),
		},

		"1N+1 ∩ 2N => 2N": {
			A: NewSubspace(1, 1),
			B: NewSubspace(2, 0),

			Expected: NewSubspace(2, 0),
		},

		"2N+1 ∩ 2N => (error: no intersection)": {
			A: NewSubspace(2, 1),
			B: NewSubspace(2, 0),

			ExpectedErr: ErrNoIntersectionSubspace,
		},

		"2N+1 ∩ 3N => 6N+3": {
			A: NewSubspace(2, 1),
			B: NewSubspace(3, 0),

			Expected: NewSubspace(6, 3),
		},

		"2N ∩ 7N+1 => 14N+8": {
			A: NewSubspace(2, 0),
			B: NewSubspace(7, 1),

			Expected: NewSubspace(14, 8),
		},
	}

	for name, cfg := range tests {
		t.Run(name, func(t *testing.T) { testFn(t, cfg) })
	}
}
