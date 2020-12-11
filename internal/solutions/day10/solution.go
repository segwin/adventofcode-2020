package day10

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/segwin/adventofcode-2020/internal/input"
)

type Solution struct{}

func (s *Solution) Run(ctx context.Context, inputFile string) {
	scanner, err := input.NewFileScanner(ctx, inputFile)
	if err != nil {
		fmt.Printf("ERROR: Failed to create input file scanner: %v\n", err)
		os.Exit(1)
	}

	adapters, err := s.getAdapters(scanner)
	if err != nil {
		fmt.Printf("ERROR: Failed to get adapters: %v\n", err)
		os.Exit(1)
	}

	s.part1(adapters)
	s.part2(adapters)
}

func (s *Solution) getAdapters(scanner input.Scanner) (adapters Adapters, err error) {
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())
		joltage, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		adapters = append(adapters, Adapter(joltage))
	}

	// sort from lowest to highest
	sort.Sort(adapters)

	// add built-in adapter (highest + 3)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	return adapters, nil
}

func (s *Solution) part1(adapters Adapters) {
	fmt.Println("\nPART 1")

	differences := map[int]int{}
	for i := 0; i < len(adapters); i++ {
		referenceJoltage := 0
		if i > 0 {
			referenceJoltage = int(adapters[i-1])
		}

		difference := int(adapters[i]) - referenceJoltage
		if difference > 3 {
			fmt.Printf("  ERROR: Found unbridgeable difference between adapters (%d jolts)\n", difference)
			os.Exit(1)
		}

		differences[difference]++
	}

	fmt.Printf("  Found %d differences\n", len(differences))

	for _, expectedDiff := range []int{1, 3} {
		if _, ok := differences[expectedDiff]; !ok {
			fmt.Printf("  ERROR: No %d jolt differences found\n", expectedDiff)
			os.Exit(1)
		}
	}

	fmt.Printf("  RESULT: Product of 1 jolt & 3 jolt differences is %d\n", differences[1]*differences[3])
}

func (s *Solution) part2(adapters Adapters) {
	fmt.Println("\nPART 2")

	adapters = append(Adapters{0}, adapters...)

	// identify all choke points in the graph, i.e. points where next is 3 away
	var chokePoints []Adapter
	for i := range adapters {
		if i == len(adapters)-1 {
			chokePoints = append(chokePoints, adapters[i])
			break
		}

		if adapters[i+1]-adapters[i] == 3 {
			// found new choke point
			chokePoints = append(chokePoints, adapters[i+1])
		}
	}

	branches := big.NewInt(1)
	for i, currentChokePoint := range chokePoints {
		lastChokePoint := Adapter(0)
		if i > 0 {
			lastChokePoint = chokePoints[i-1]
		}

		var relativeAdapters Adapters
		for _, adapter := range adapters {
			if adapter == lastChokePoint {
				relativeAdapters = append(relativeAdapters, 0)
				continue
			}

			if len(relativeAdapters) > 0 {
				relativeAdapters = append(relativeAdapters, adapter-lastChokePoint)
				if adapter == currentChokePoint {
					break
				}
			}
		}

		idealBranches := s.calculateBranches(int(currentChokePoint - lastChokePoint)) // start with ideal number (full tree)
		holes := s.getHoles(relativeAdapters)                                         // remove branches for each hole in this ideal tree

		// compute total number of missing branches, starting with highest hole value
		totalRemoved := big.NewInt(0)
		for j := len(holes) - 1; j >= 0; j-- {
			// remove potential branches for this hole, subtracting branches previously
			// counted for higher hole values
			//
			// R0 = (potential branches for highest hole value)
			//    = n(hole) * n(max - hole)
			//
			// R1 = (potential branches for next highest hole value)
			//    = n(hole) * n(max - hole) - R0
			//
			// R2 = (potential branches for next hole value)
			//    = n(hole) * n(max - hole) - R1 + R0
			//
			// ...
			//
			// Rn = (potential branches for Nth highest hole)
			//    = n(hole) * n(max - hole) - Rn_1 + Rn_2 - Rn_3 + ...

			branchesPerHole := s.branchesPerNode(int(currentChokePoint-lastChokePoint), holes[j])
			numHoles := s.numNodes(holes[j], 0)
			removed := big.NewInt(0).Mul(numHoles, branchesPerHole)

			if j < len(holes)-1 {
				numChildBranches := s.countBranches(int(currentChokePoint-lastChokePoint), holes[j+1:], holes[j])
				removed.Sub(removed, numChildBranches.Mul(numChildBranches, numHoles))
			}

			totalRemoved = big.NewInt(0).Add(totalRemoved, removed)
		}

		branches.Mul(branches, big.NewInt(0).Sub(idealBranches, totalRemoved))
	}

	fmt.Printf("  RESULT: Found %s unique branches\n", branches.Text(10))
}

func (s *Solution) getHoles(adapters Adapters) (holes []int) {
	max := int(adapters[len(adapters)-1])

	offset := 0
	for i := 0; i < max; i++ {
		if actual := int(adapters[i+offset]); actual == i {
			continue
		}

		offset--
		holes = append(holes, i)
	}

	return holes
}

func (s *Solution) calculateBranches(max int) *big.Int {
	prevN := [3]*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(2)}

	for i := 2; i < max; i++ {
		n0 := prevN[1]
		n1 := prevN[2]

		n2 := big.NewInt(0).Add(prevN[2], prevN[1])
		n2 = n2.Add(n2, prevN[0])

		prevN[0], prevN[1], prevN[2] = n0, n1, n2
	}

	if max < 2 {
		return prevN[max]
	}

	return prevN[2]
}

func (s *Solution) numNodes(current, parent int) *big.Int {
	return s.calculateBranches(current - parent)
}

func (s *Solution) branchesPerNode(max, node int) *big.Int {
	return s.calculateBranches(max - node)
}

func (s *Solution) countBranches(max int, nodes []int, parent int) *big.Int {
	count := big.NewInt(0)

	for i, node := range nodes {
		numNodes := s.numNodes(node, parent)
		branchesPerNode := s.branchesPerNode(max, node)
		count.Add(count, big.NewInt(0).Mul(numNodes, branchesPerNode))

		// recurse to subtract counts included in sibling nodes
		if i < len(nodes)-1 {
			count.Sub(count, s.countBranches(max, nodes[i+1:], nodes[i]))
		}
	}

	return count
}
