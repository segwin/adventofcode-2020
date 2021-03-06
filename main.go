package main

import (
	"fmt"
	"os"

	"github.com/segwin/adventofcode-2020/cmd"
)

func main() {
	if err := cmd.New("aoc").Execute(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
