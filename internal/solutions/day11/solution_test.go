package day11

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/segwin/adventofcode-2020/internal/input"
)

func TestGetLayout(t *testing.T) {
	inputStr := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`

	expected := Layout{
		[]Seat{Empty, Floor, Empty, Empty, Floor, Empty, Empty, Floor, Empty, Empty},
		[]Seat{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Floor, Empty, Empty},
		[]Seat{Empty, Floor, Empty, Floor, Empty, Floor, Floor, Empty, Floor, Floor},
		[]Seat{Empty, Empty, Empty, Empty, Floor, Empty, Empty, Floor, Empty, Empty},
		[]Seat{Empty, Floor, Empty, Empty, Floor, Empty, Empty, Floor, Empty, Empty},
		[]Seat{Empty, Floor, Empty, Empty, Empty, Empty, Empty, Floor, Empty, Empty},
		[]Seat{Floor, Floor, Empty, Floor, Empty, Floor, Floor, Floor, Floor, Floor},
		[]Seat{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		[]Seat{Empty, Floor, Empty, Empty, Empty, Empty, Empty, Empty, Floor, Empty},
		[]Seat{Empty, Floor, Empty, Empty, Empty, Empty, Empty, Floor, Empty, Empty},
	}

	var s Solution
	scanner := input.NewStringScanner(context.Background(), inputStr)

	got, err := s.getLayout(scanner)
	if err != nil {
		t.Errorf("Got %v, expected nil", err)
	}

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("Unexpected diff:\n%v", diff)
	}
}

func TestSolutionPart1(t *testing.T) {
	input := Layout{
		[]Seat{Empty, Floor, Empty, Empty, Floor, Empty, Empty, Floor, Empty, Empty},
		[]Seat{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Floor, Empty, Empty},
		[]Seat{Empty, Floor, Empty, Floor, Empty, Floor, Floor, Empty, Floor, Floor},
		[]Seat{Empty, Empty, Empty, Empty, Floor, Empty, Empty, Floor, Empty, Empty},
		[]Seat{Empty, Floor, Empty, Empty, Floor, Empty, Empty, Floor, Empty, Empty},
		[]Seat{Empty, Floor, Empty, Empty, Empty, Empty, Empty, Floor, Empty, Empty},
		[]Seat{Floor, Floor, Empty, Floor, Empty, Floor, Floor, Floor, Floor, Floor},
		[]Seat{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		[]Seat{Empty, Floor, Empty, Empty, Empty, Empty, Empty, Empty, Floor, Empty},
		[]Seat{Empty, Floor, Empty, Empty, Empty, Empty, Empty, Floor, Empty, Empty},
	}

	expectedStages := []Layout{
		{
			[]Seat{Occupied, Floor, Occupied, Occupied, Floor, Occupied, Occupied, Floor, Occupied, Occupied},
			[]Seat{Occupied, Occupied, Occupied, Occupied, Occupied, Occupied, Occupied, Floor, Occupied, Occupied},
			[]Seat{Occupied, Floor, Occupied, Floor, Occupied, Floor, Floor, Occupied, Floor, Floor},
			[]Seat{Occupied, Occupied, Occupied, Occupied, Floor, Occupied, Occupied, Floor, Occupied, Occupied},
			[]Seat{Occupied, Floor, Occupied, Occupied, Floor, Occupied, Occupied, Floor, Occupied, Occupied},
			[]Seat{Occupied, Floor, Occupied, Occupied, Occupied, Occupied, Occupied, Floor, Occupied, Occupied},
			[]Seat{Floor, Floor, Occupied, Floor, Occupied, Floor, Floor, Floor, Floor, Floor},
			[]Seat{Occupied, Occupied, Occupied, Occupied, Occupied, Occupied, Occupied, Occupied, Occupied, Occupied},
			[]Seat{Occupied, Floor, Occupied, Occupied, Occupied, Occupied, Occupied, Occupied, Floor, Occupied},
			[]Seat{Occupied, Floor, Occupied, Occupied, Occupied, Occupied, Occupied, Floor, Occupied, Occupied},
		},

		{
			[]Seat{Occupied, Floor, Empty, Empty, Floor, Empty, Occupied, Floor, Occupied, Occupied},
			[]Seat{Occupied, Empty, Empty, Empty, Empty, Empty, Empty, Floor, Empty, Occupied},
			[]Seat{Empty, Floor, Empty, Floor, Empty, Floor, Floor, Empty, Floor, Floor},
			[]Seat{Occupied, Empty, Empty, Empty, Floor, Empty, Empty, Floor, Empty, Occupied},
			[]Seat{Occupied, Floor, Empty, Empty, Floor, Empty, Empty, Floor, Empty, Empty},
			[]Seat{Occupied, Floor, Empty, Empty, Empty, Empty, Occupied, Floor, Occupied, Occupied},
			[]Seat{Floor, Floor, Empty, Floor, Empty, Floor, Floor, Floor, Floor, Floor},
			[]Seat{Occupied, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Occupied},
			[]Seat{Occupied, Floor, Empty, Empty, Empty, Empty, Empty, Empty, Floor, Empty},
			[]Seat{Occupied, Floor, Occupied, Empty, Empty, Empty, Empty, Floor, Occupied, Occupied},
		},

		{
			[]Seat{Occupied, Floor, Occupied, Occupied, Floor, Empty, Occupied, Floor, Occupied, Occupied},
			[]Seat{Occupied, Empty, Occupied, Occupied, Occupied, Empty, Empty, Floor, Empty, Occupied},
			[]Seat{Empty, Floor, Occupied, Floor, Occupied, Floor, Floor, Occupied, Floor, Floor},
			[]Seat{Occupied, Empty, Occupied, Occupied, Floor, Occupied, Occupied, Floor, Empty, Occupied},
			[]Seat{Occupied, Floor, Occupied, Occupied, Floor, Empty, Empty, Floor, Empty, Empty},
			[]Seat{Occupied, Floor, Occupied, Occupied, Occupied, Empty, Occupied, Floor, Occupied, Occupied},
			[]Seat{Floor, Floor, Occupied, Floor, Occupied, Floor, Floor, Floor, Floor, Floor},
			[]Seat{Occupied, Empty, Occupied, Occupied, Occupied, Occupied, Occupied, Occupied, Empty, Occupied},
			[]Seat{Occupied, Floor, Empty, Empty, Occupied, Occupied, Occupied, Empty, Floor, Empty},
			[]Seat{Occupied, Floor, Occupied, Empty, Occupied, Occupied, Occupied, Floor, Occupied, Occupied},
		},

		{
			[]Seat{Occupied, Floor, Occupied, Empty, Floor, Empty, Occupied, Floor, Occupied, Occupied},
			[]Seat{Occupied, Empty, Empty, Empty, Occupied, Empty, Empty, Floor, Empty, Occupied},
			[]Seat{Empty, Floor, Empty, Floor, Empty, Floor, Floor, Occupied, Floor, Floor},
			[]Seat{Occupied, Empty, Empty, Empty, Floor, Occupied, Occupied, Floor, Empty, Occupied},
			[]Seat{Occupied, Floor, Empty, Empty, Floor, Empty, Empty, Floor, Empty, Empty},
			[]Seat{Occupied, Floor, Empty, Empty, Occupied, Empty, Occupied, Floor, Occupied, Occupied},
			[]Seat{Floor, Floor, Empty, Floor, Empty, Floor, Floor, Floor, Floor, Floor},
			[]Seat{Occupied, Empty, Occupied, Empty, Empty, Empty, Empty, Occupied, Empty, Occupied},
			[]Seat{Occupied, Floor, Empty, Empty, Empty, Empty, Empty, Empty, Floor, Empty},
			[]Seat{Occupied, Floor, Occupied, Empty, Occupied, Empty, Occupied, Floor, Occupied, Occupied},
		},

		{
			[]Seat{Occupied, Floor, Occupied, Empty, Floor, Empty, Occupied, Floor, Occupied, Occupied},
			[]Seat{Occupied, Empty, Empty, Empty, Occupied, Empty, Empty, Floor, Empty, Occupied},
			[]Seat{Empty, Floor, Occupied, Floor, Empty, Floor, Floor, Occupied, Floor, Floor},
			[]Seat{Occupied, Empty, Occupied, Occupied, Floor, Occupied, Occupied, Floor, Empty, Occupied},
			[]Seat{Occupied, Floor, Occupied, Empty, Floor, Empty, Empty, Floor, Empty, Empty},
			[]Seat{Occupied, Floor, Occupied, Empty, Occupied, Empty, Occupied, Floor, Occupied, Occupied},
			[]Seat{Floor, Floor, Empty, Floor, Empty, Floor, Floor, Floor, Floor, Floor},
			[]Seat{Occupied, Empty, Occupied, Empty, Occupied, Occupied, Empty, Occupied, Empty, Occupied},
			[]Seat{Occupied, Floor, Empty, Empty, Empty, Empty, Empty, Empty, Floor, Empty},
			[]Seat{Occupied, Floor, Occupied, Empty, Occupied, Empty, Occupied, Floor, Occupied, Occupied},
		},

		{
			[]Seat{Occupied, Floor, Occupied, Empty, Floor, Empty, Occupied, Floor, Occupied, Occupied},
			[]Seat{Occupied, Empty, Empty, Empty, Occupied, Empty, Empty, Floor, Empty, Occupied},
			[]Seat{Empty, Floor, Occupied, Floor, Empty, Floor, Floor, Occupied, Floor, Floor},
			[]Seat{Occupied, Empty, Occupied, Occupied, Floor, Occupied, Occupied, Floor, Empty, Occupied},
			[]Seat{Occupied, Floor, Occupied, Empty, Floor, Empty, Empty, Floor, Empty, Empty},
			[]Seat{Occupied, Floor, Occupied, Empty, Occupied, Empty, Occupied, Floor, Occupied, Occupied},
			[]Seat{Floor, Floor, Empty, Floor, Empty, Floor, Floor, Floor, Floor, Floor},
			[]Seat{Occupied, Empty, Occupied, Empty, Occupied, Occupied, Empty, Occupied, Empty, Occupied},
			[]Seat{Occupied, Floor, Empty, Empty, Empty, Empty, Empty, Empty, Floor, Empty},
			[]Seat{Occupied, Floor, Occupied, Empty, Occupied, Empty, Occupied, Floor, Occupied, Occupied},
		},
	}

	var s Solution

	prevLayout := input
	for i, expectedStage := range expectedStages {
		got := s.evolveLayout(prevLayout, 4, false)
		if diff := cmp.Diff(expectedStage, got); diff != "" {
			t.Errorf("Unexpected diff (iteration: %d):\n%v", i+1, diff)
		}

		if diff := cmp.Diff(prevLayout, got); diff == "" && i < len(expectedStages)-1 {
			t.Error("Expected diff, got \"\"")
		}

		if got.Equals(prevLayout) && i < len(expectedStages)-1 {
			t.Errorf("Reached equilibrium sooner than expected: got %d, expected %d", i+1, len(expectedStages))
		}

		prevLayout = got
	}

	if got, expected := prevLayout.Count(Occupied), 37; got != expected {
		t.Errorf("Got %v, expected %v", got, expected)
	}
}
