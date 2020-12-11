package day10

import (
	"sort"
	"strconv"
)

// Adapter represents a joltage adapter.
type Adapter int

// Adapters is a sortable list of Adapter objects. It also implements a hash function
// so that it can be used as a map key.
type Adapters []Adapter

var _ sort.Interface = (Adapters)(nil)

func (a Adapters) Len() int {
	return len(a)
}

func (a Adapters) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a Adapters) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Adapters) Hash() (hash string) {
	for _, adapter := range a {
		hash += strconv.Itoa(int(adapter)) + "-"
	}

	return hash
}
