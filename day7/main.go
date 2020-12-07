package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getLines(scanner *bufio.Scanner, callback func(line string)) {
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("File read error: %v\n", err)
			os.Exit(1)
		}

		line := strings.Trim(scanner.Text(), " \t\v\n")
		callback(line)
	}
}

func part1(colour string, bags map[string]Bag) {
	count := 0
	for _, bag := range bags {
		if bag.Contains(colour, bags) {
			count++
		}
	}

	// print result
	fmt.Printf("PART 1 RESULT: Found %d bags that can contain a %s bag\n", count, colour)
}

func part2(colour string, bags map[string]Bag) {
	bag, ok := bags[colour]
	if !ok {
		fmt.Printf("PART 2 ERROR: Failed to find %s bag\n", colour)
		os.Exit(1)
	}

	count := int64(0)
	subBags := bag.SubBags(bags)
	for _, subBagCount := range subBags {
		count += subBagCount
	}

	// print result
	fmt.Printf("PART 1 RESULT: Found total of %d bags inside the %s bag\n", count, colour)
}

func main() {
	file, err := os.Open("day7/input")
	if err != nil {
		fmt.Printf("File open error: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	bags := map[string]Bag{}
	callback := func(line string) {
		newBag := &bag{}
		if err := newBag.Unmarshal(line); err != nil {
			fmt.Printf("Failed to unmarshal response for line %q: %v\n", line, err)
			os.Exit(1)
		}

		bags[newBag.Colour()] = newBag
	}

	getLines(bufio.NewScanner(file), callback)

	part1("shiny gold", bags)
	part2("shiny gold", bags)
}
