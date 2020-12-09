package day6

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnmarshalResponses(t *testing.T) {
	t.Parallel()

	type Test struct {
		// inputs
		encoded   []string
		responses Responses // used to provide different Responses implementations

		// outputs
		expectedResponses Responses
		expectedYesCount  int
		expectedErr       error
	}

	testFn := func(t *testing.T, cfg Test) {
		t.Parallel()

		var err error
		for _, line := range cfg.encoded {
			err = cfg.responses.UnmarshalNew(line)
			if err != nil {
				break
			}
		}

		if !errors.Is(err, cfg.expectedErr) {
			t.Errorf("Got %v, expected %v", err, cfg.expectedErr)
		}

		if err != nil {
			return // we're done
		}

		if diff := cmp.Diff(cfg.expectedResponses, cfg.responses); diff != "" {
			t.Errorf("Unexpected diff:\n%v", diff)
		}

		if got, expected := cfg.responses.YesCount(), cfg.expectedYesCount; got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	}

	tests := map[string]Test{
		"error: non-latin alphabetical character": {
			encoded:     []string{"aèc"},
			responses:   NewIndividualResponses(),
			expectedErr: ErrInvalidChar,
		},

		"error: numeric character": {
			encoded:     []string{"a1c"},
			responses:   NewIndividualResponses(),
			expectedErr: ErrInvalidChar,
		},

		"ok: \"abc\"": {
			encoded:           []string{"abc"},
			responses:         NewIndividualResponses(),
			expectedResponses: individualResponses{'a': true, 'b': true, 'c': true},
			expectedYesCount:  3,
		},

		"ok: empty input": {
			encoded:           []string{""},
			responses:         NewIndividualResponses(),
			expectedResponses: NewIndividualResponses(),
			expectedYesCount:  0,
		},

		"ok: single letter": {
			encoded:           []string{"a"},
			responses:         NewIndividualResponses(),
			expectedResponses: individualResponses{'a': true},
			expectedYesCount:  1,
		},

		"ok: whole alphabet": {
			encoded:           []string{"abcdefghijklmnopqrstuvwxyz"},
			responses:         NewIndividualResponses(),
			expectedResponses: individualResponses{'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true, 'g': true, 'h': true, 'i': true, 'j': true, 'k': true, 'l': true, 'm': true, 'n': true, 'o': true, 'p': true, 'q': true, 'r': true, 's': true, 't': true, 'u': true, 'v': true, 'w': true, 'x': true, 'y': true, 'z': true},
			expectedYesCount:  26,
		},

		"ok: duplicate characters": {
			encoded:           []string{"abbcdd"},
			responses:         NewIndividualResponses(),
			expectedResponses: individualResponses{'a': true, 'b': true, 'c': true, 'd': true},
			expectedYesCount:  4,
		},

		"ok: multiple lines, no duplicates": {
			encoded:           []string{"ab", "cd"},
			responses:         NewIndividualResponses(),
			expectedResponses: individualResponses{'a': true, 'b': true, 'c': true, 'd': true},
			expectedYesCount:  4,
		},

		"ok: multiple lines with duplicates across lines": {
			encoded:           []string{"abc", "bcd"},
			responses:         NewIndividualResponses(),
			expectedResponses: individualResponses{'a': true, 'b': true, 'c': true, 'd': true},
			expectedYesCount:  4,
		},

		"ok: multiple lines with duplicates across lines, unanimous responses": {
			encoded:   []string{"abc", "bcd"},
			responses: NewUnanimousResponses(),
			expectedResponses: &unanimousResponses{
				RespondeeCount: 2,
				Affirmatives:   map[rune][]bool{'a': {true}, 'b': {true, true}, 'c': {true, true}, 'd': {true}},
			},
			expectedYesCount: 2,
		},
	}

	for name, cfg := range tests {
		t.Run(name, func(t *testing.T) { testFn(t, cfg) })
	}
}
