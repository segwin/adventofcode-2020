package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrBadKVPair = errors.New("invalid key-value pair")
)

type FieldValidator func(value string) (ok bool)

var (
	allowedECLValues = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	keyValidation = map[string]FieldValidator{
		// Birth Year
		"byr": func(value string) (ok bool) {
			year, ok := ToYear(value)
			if !ok {
				return false
			}

			return year >= 1920 && year <= 2002
		},

		// Issue Year
		"iyr": func(value string) (ok bool) {
			year, ok := ToYear(value)
			if !ok {
				return false
			}

			return year >= 2010 && year <= 2020
		},

		// Expiration Year
		"eyr": func(value string) (ok bool) {
			year, ok := ToYear(value)
			if !ok {
				return false
			}

			return year >= 2020 && year <= 2030
		},

		// Height
		"hgt": func(value string) (ok bool) {
			var min, max int
			if strings.HasSuffix(value, "cm") {
				min = 150
				max = 193
			} else if strings.HasSuffix(value, "in") {
				min = 59
				max = 76
			} else {
				return false
			}

			for _, unit := range []string{"cm", "in"} {
				value = strings.TrimSuffix(value, unit)
			}

			height, err := strconv.Atoi(value)
			if err != nil {
				return false
			}

			return height >= min && height <= max
		},

		// Hair Color
		"hcl": func(value string) (ok bool) {
			if !strings.HasPrefix(value, "#") {
				return false
			}

			hexColour := strings.ToLower(strings.TrimPrefix(value, "#"))
			if len(hexColour) != 6 {
				return false
			}

			for _, character := range hexColour {
				if !unicode.In(character, unicode.Hex_Digit) {
					return false
				}
			}

			return true
		},

		// Eye Color
		"ecl": func(value string) (ok bool) {
			for _, allowedValue := range allowedECLValues {
				if value == allowedValue {
					return true
				}
			}

			return false
		},

		// Passport ID
		"pid": func(value string) (ok bool) {
			if len(value) != 9 {
				return false
			}

			_, err := strconv.Atoi(value)
			return err == nil
		},

		// Country ID - disabled to let me board the plane
		//"cid": {},
	}
)

func ToYear(value string) (year int, ok bool) {
	if len(value) != 4 {
		return 0, false
	}

	year, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}

	return year, true
}

type Passport map[string]string

func (p Passport) Unmarshal(lines []string) error {
	data := strings.Join(lines, " ")

	kvPairs := strings.Split(strings.Trim(data, " "), " ")
	for _, kvPair := range kvPairs {
		fields := strings.Split(kvPair, ":")
		if len(fields) != 2 {
			return fmt.Errorf("%w (%q)", ErrBadKVPair, kvPair)
		}

		p[fields[0]] = fields[1]
	}

	return nil
}

func (p Passport) IsValid() bool {
	for key, isValid := range keyValidation {
		if value, ok := p[key]; !ok || !isValid(value) {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("day4/input")
	if err != nil {
		fmt.Printf("File open error: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	// let's use a streaming approach for a change instead of loading everything into memory :)
	scanner := bufio.NewScanner(file)

	// read all lines, collecting passport lines along the way & unmarshaling when end-of-passport reached
	var passportLines []string
	var validPassports []Passport

	parseAndAddValid := func() {
		passport := Passport{}
		if err := passport.Unmarshal(passportLines); err != nil {
			fmt.Printf("Fatal error: %s\n", err)
			os.Exit(1)
		}

		if passport.IsValid() {
			validPassports = append(validPassports, passport)
		}

		passportLines = nil
	}

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("File read error: %v\n", err)
			os.Exit(1)
		}

		line := strings.Trim(scanner.Text(), " \t\v\n")
		if len(line) > 0 {
			passportLines = append(passportLines, line)
			continue
		}

		// reached end of passport
		parseAndAddValid()
	}

	if len(passportLines) > 0 {
		parseAndAddValid()
	}

	fmt.Printf("RESULT: Found %d valid passports\n", len(validPassports))
}
