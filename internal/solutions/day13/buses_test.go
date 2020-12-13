package day13

import (
	"errors"
	"testing"
)

func TestScheduleIntersect(t *testing.T) {
	type Test struct {
		Schedule Schedule

		Expected    int64
		ExpectedErr error
	}

	testFn := func(t *testing.T, cfg Test) {
		lowestTime, err := cfg.Schedule.FindLowestTime()
		if got, expected := err, cfg.ExpectedErr; !errors.Is(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}

		if got, expected := lowestTime, cfg.Expected; got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	}

	tests := map[string]Test{
		"17,x,13,19": {
			Schedule: Schedule{
				Buses: []*Subspace{
					NewSubspace(17, 0),
					NewSubspace(13, 2),
					NewSubspace(19, 3),
				},
			},

			Expected: 3417,
		},

		"67,7,59,61": {
			Schedule: Schedule{
				Buses: []*Subspace{
					NewSubspace(67, 0),
					NewSubspace(7, 1),
					NewSubspace(59, 2),
					NewSubspace(61, 3),
				},
			},

			Expected: 754018,
		},

		"67,x,7,59,61": {
			Schedule: Schedule{
				Buses: []*Subspace{
					NewSubspace(67, 0),
					NewSubspace(7, 2),
					NewSubspace(59, 3),
					NewSubspace(61, 4),
				},
			},

			Expected: 779210,
		},

		"67,7,x,59,61": {
			Schedule: Schedule{
				Buses: []*Subspace{
					NewSubspace(67, 0),
					NewSubspace(7, 1),
					NewSubspace(59, 3),
					NewSubspace(61, 4),
				},
			},

			Expected: 1261476,
		},

		"1789,37,47,1889": {
			Schedule: Schedule{
				Buses: []*Subspace{
					NewSubspace(1789, 0),
					NewSubspace(37, 1),
					NewSubspace(47, 2),
					NewSubspace(1889, 3),
				},
			},

			Expected: 1202161486,
		},

		"7,13,x,x,59,x,31,19": {
			Schedule: Schedule{
				Buses: []*Subspace{
					NewSubspace(7, 0),
					NewSubspace(13, 1),
					NewSubspace(59, 4),
					NewSubspace(31, 6),
					NewSubspace(19, 7),
				},
			},

			Expected: 1068781,
		},

		"9,x,x,15": {
			Schedule: Schedule{
				Buses: []*Subspace{
					NewSubspace(9, 0),
					NewSubspace(15, -3%38),
				},
			},

			Expected: 18,
		},

		"9,x,x,x,x,x,15": {
			Schedule: Schedule{
				Buses: []*Subspace{
					NewSubspace(30, 0),
					NewSubspace(38, -6%38),
				},
			},

			Expected: 120,
		},

		"9,x,x,x,x,12": {
			Schedule: Schedule{
				Buses: []*Subspace{
					NewSubspace(12, 0),
					NewSubspace(9, -5%9),
				},
			},

			ExpectedErr: ErrNoIntersectionSubspace,
		},
	}

	for name, cfg := range tests {
		t.Run(name, func(t *testing.T) { testFn(t, cfg) })
	}
}
