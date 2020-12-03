package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	ErrBadEntryLine = errors.New("failed to unmarshal password entry line")
	ErrBadPolicy    = errors.New("failed to unmarshal password policy")
)

type PasswordPolicy interface {
	Unmarshal(policyStr string) error
	Validate(password string) (ok bool)
}

type oldPasswordPolicy struct {
	RequiredLetter byte
	Min            int
	Max            int
}

func (p *oldPasswordPolicy) Unmarshal(policyStr string) error {
	policyComponents := strings.Split(policyStr, " ")
	if len(policyComponents) != 2 {
		return fmt.Errorf("%w: got %d overall components, expected 2", ErrBadPolicy, len(policyComponents))
	}

	policyCountComponents := strings.Split(policyComponents[0], "-")
	if len(policyCountComponents) != 2 {
		return fmt.Errorf("%w: got %d count components, expected 2", ErrBadPolicy, len(policyCountComponents))
	}

	min, err := strconv.Atoi(policyCountComponents[0])
	if err != nil {
		return fmt.Errorf("%w: invalid min count (%q)", ErrBadPolicy, err)
	}

	max, err := strconv.Atoi(policyCountComponents[1])
	if err != nil {
		return fmt.Errorf("%w: invalid max count (%q)", ErrBadPolicy, err)
	}

	if len(policyComponents[1]) != 1 {
		return fmt.Errorf("%w: got %d required letters, expected 1", ErrBadPolicy, len(policyComponents[1]))
	}

	p.RequiredLetter = policyComponents[1][0]
	p.Min = min
	p.Max = max
	return nil
}

func (p *oldPasswordPolicy) Validate(password string) (ok bool) {
	letterCount := 0
	for i := range password {
		if password[i] == p.RequiredLetter {
			letterCount++
		}
	}

	return letterCount >= p.Min && letterCount <= p.Max
}

type passwordPolicy struct {
	RequiredLetter byte
	Positions      [2]int
}

func (p *passwordPolicy) Unmarshal(policyStr string) error {
	policyComponents := strings.Split(policyStr, " ")
	if len(policyComponents) != 2 {
		return fmt.Errorf("%w: got %d overall components, expected 2", ErrBadPolicy, len(policyComponents))
	}

	countComponents := strings.Split(policyComponents[0], "-")
	if len(countComponents) != len(p.Positions) {
		return fmt.Errorf("%w: got %d count components, expected %d", ErrBadPolicy, len(countComponents), len(p.Positions))
	}

	positions := make([]int, len(p.Positions))
	for i, c := range countComponents {
		oneIndexedPos, err := strconv.Atoi(c)
		if err != nil {
			return fmt.Errorf("%w: invalid position %d value (%q)", ErrBadPolicy, i, err)
		}

		positions[i] = oneIndexedPos - 1
	}

	if len(policyComponents[1]) != 1 {
		return fmt.Errorf("%w: got %d required letters, expected 1", ErrBadPolicy, len(policyComponents[1]))
	}

	p.RequiredLetter = policyComponents[1][0]
	copy(p.Positions[:], positions)

	return nil
}

func (p *passwordPolicy) Validate(password string) (ok bool) {
	foundPositions := make(map[int]bool)
	for i := range password {
		if password[i] == p.RequiredLetter {
			foundPositions[i] = true
		}
	}

	foundOne := false
	for _, position := range p.Positions {
		if foundPositions[position] {
			if foundOne {
				return false // letter found at >1 position
			}

			foundOne = true
		}
	}

	return foundOne
}

type PasswordEntry struct {
	Policy   PasswordPolicy
	Password string
}

func (e *PasswordEntry) IsValid() bool {
	return e.Policy.Validate(e.Password)
}

func UnmarshalEntry(line string, oldPolicy bool) (*PasswordEntry, error) {
	components := strings.Split(line, ":")
	if len(components) != 2 {
		return nil, fmt.Errorf("%w: got %d global components, expected 2", ErrBadEntryLine, len(components))
	}

	var policy PasswordPolicy = &passwordPolicy{}
	if oldPolicy {
		policy = &oldPasswordPolicy{}
	}

	if err := policy.Unmarshal(strings.Trim(components[0], " \t\n\v")); err != nil {
		return nil, err
	}

	password := strings.Trim(components[1], " \t\v\n")

	return &PasswordEntry{Policy: policy, Password: password}, nil
}

func main() {
	f, err := ioutil.ReadFile("day2/input")
	if err != nil {
		fmt.Printf("Read error\n")
		panic(err)
	}

	lines := strings.Split(strings.Trim(string(f), " \t\v\n"), "\n")
	entries := make([]*PasswordEntry, 0, len(lines))

	for i, line := range lines {
		entry, err := UnmarshalEntry(line, false)
		if err != nil {
			fmt.Printf("Parse error on line %d: %q\n", i, err)
			continue
		}

		entries = append(entries, entry)
	}

	validCount := 0
	for _, entry := range entries {
		if entry.IsValid() {
			validCount++
		}
	}

	fmt.Printf("RESULT: Got %d valid lines\n", validCount)
}
