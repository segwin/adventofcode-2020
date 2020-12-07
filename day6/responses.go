package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidChar   = errors.New("invalid character in responses line")
	ErrDuplicateChar = errors.New("duplicate char in response line")
)

type Responses interface {
	UnmarshalNew(responses string) error
	YesCount() int
}

// unanimousResponses implements the logic specified in part 1, i.e. questions that
// *anyone* answered yes to are counted as a yes.
type individualResponses map[rune]bool

func NewIndividualResponses() Responses {
	return individualResponses{}
}

func (r individualResponses) UnmarshalNew(responses string) error {
	responses = strings.ToLower(responses) // normalise line
	for _, letter := range responses {
		if !isLatinAlpha(letter) {
			return fmt.Errorf("%w (%s)", ErrInvalidChar, string(letter))
		}

		r[letter] = true
	}

	return nil
}

func (r individualResponses) YesCount() (count int) {
	for _, yes := range r {
		if yes {
			count++
		}
	}

	return count
}

// unanimousResponses implements the logic specified in part 2, i.e. only questions
// that *everyone* responded "yes" to are counted as a yes.
type unanimousResponses struct {
	RespondeeCount int
	Affirmatives   map[rune][]bool
}

func NewUnanimousResponses() Responses {
	return &unanimousResponses{Affirmatives: map[rune][]bool{}}
}

func (r *unanimousResponses) UnmarshalNew(responses string) error {
	r.RespondeeCount++

	responses = strings.ToLower(responses) // normalise line
	for _, letter := range responses {
		if !isLatinAlpha(letter) {
			return fmt.Errorf("%w (%s)", ErrInvalidChar, string(letter))
		}

		if affirmatives, ok := r.Affirmatives[letter]; ok && len(affirmatives) >= r.RespondeeCount {
			// already have an affirmative for this line, meaning this is a duplicate on the same line
			return fmt.Errorf("%w (%s)", ErrDuplicateChar, string(letter))
		}

		r.Affirmatives[letter] = append(r.Affirmatives[letter], true)
	}

	return nil
}

func (r *unanimousResponses) YesCount() (count int) {
	for _, affirmatives := range r.Affirmatives {
		if len(affirmatives) == r.RespondeeCount {
			count++ // all respondees responded affirmatively to this question
		}
	}

	return count
}

func isLatinAlpha(character rune) bool {
	return character >= 'a' && character <= 'z'
}
