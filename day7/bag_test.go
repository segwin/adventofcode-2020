package main

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBagUnmarshal(t *testing.T) {
	t.Parallel()

	type Test struct {
		// inputs
		msg string

		// outputs
		expected    Bag
		expectedErr error
	}

	testFn := func(t *testing.T, cfg Test) {
		t.Parallel()

		var myBag Bag = &bag{}

		err := myBag.Unmarshal(cfg.msg)
		if !errors.Is(err, cfg.expectedErr) {
			t.Errorf("Got %v, expected %v", err, cfg.expectedErr)
		}

		if err != nil {
			return // we're done
		}

		if diff := cmp.Diff(cfg.expected, myBag, cmp.AllowUnexported(bag{})); diff != "" {
			t.Errorf("Unexpected diff:\n%v", diff)
		}
	}

	tests := map[string]Test{
		"ok: single bag containing a shiny gold bag": {
			msg:      "bright white bags contain 1 shiny gold bag.",
			expected: &bag{colour: "bright white", contains: map[string]int64{"shiny gold": 1}},
		},
	}

	for name, cfg := range tests {
		t.Run(name, func(t *testing.T) { testFn(t, cfg) })
	}
}

func TestBagCanContain(t *testing.T) {
	t.Parallel()

	type Test struct {
		// inputs
		bag         Bag
		countColour string
		allBags     map[string]Bag

		// outputs
		expected bool
	}

	testFn := func(t *testing.T, cfg Test) {
		t.Parallel()

		if got, expected := cfg.bag.Contains(cfg.countColour, cfg.allBags), cfg.expected; got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	}

	tests := map[string]Test{
		"false: single bag containing a vibrant plum": {
			bag:         &bag{colour: "bright white", contains: map[string]int64{"vibrant plum": 1}},
			countColour: "shiny gold",
			allBags:     map[string]Bag{},

			expected: false,
		},

		"true: single bag containing a shiny gold bag": {
			bag:         &bag{colour: "bright white", contains: map[string]int64{"shiny gold": 1}},
			countColour: "shiny gold",
			allBags:     map[string]Bag{},

			expected: true,
		},
	}

	for name, cfg := range tests {
		t.Run(name, func(t *testing.T) { testFn(t, cfg) })
	}
}
