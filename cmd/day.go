package cmd

import (
	"fmt"

	"github.com/segwin/adventofcode-2020/internal/solutions"
	"github.com/segwin/adventofcode-2020/internal/solutions/day1"
	"github.com/segwin/adventofcode-2020/internal/solutions/day10"
	"github.com/segwin/adventofcode-2020/internal/solutions/day11"
	"github.com/segwin/adventofcode-2020/internal/solutions/day12"
	"github.com/segwin/adventofcode-2020/internal/solutions/day13"
	"github.com/segwin/adventofcode-2020/internal/solutions/day14"
	"github.com/segwin/adventofcode-2020/internal/solutions/day15"
	"github.com/segwin/adventofcode-2020/internal/solutions/day16"
	"github.com/segwin/adventofcode-2020/internal/solutions/day17"
	"github.com/segwin/adventofcode-2020/internal/solutions/day18"
	"github.com/segwin/adventofcode-2020/internal/solutions/day19"
	"github.com/segwin/adventofcode-2020/internal/solutions/day2"
	"github.com/segwin/adventofcode-2020/internal/solutions/day20"
	"github.com/segwin/adventofcode-2020/internal/solutions/day21"
	"github.com/segwin/adventofcode-2020/internal/solutions/day22"
	"github.com/segwin/adventofcode-2020/internal/solutions/day23"
	"github.com/segwin/adventofcode-2020/internal/solutions/day24"
	"github.com/segwin/adventofcode-2020/internal/solutions/day25"
	"github.com/segwin/adventofcode-2020/internal/solutions/day3"
	"github.com/segwin/adventofcode-2020/internal/solutions/day4"
	"github.com/segwin/adventofcode-2020/internal/solutions/day5"
	"github.com/segwin/adventofcode-2020/internal/solutions/day6"
	"github.com/segwin/adventofcode-2020/internal/solutions/day7"
	"github.com/segwin/adventofcode-2020/internal/solutions/day8"
	"github.com/segwin/adventofcode-2020/internal/solutions/day9"
	"github.com/spf13/cobra"
)

var (
	inputFiles = map[int]*string{}
)

func newDayCommand(day int, solution solutions.Solution) *cobra.Command {
	dayCmd := &cobra.Command{
		Use:   fmt.Sprintf("day%d", day),
		Short: fmt.Sprintf("Run the solution for day %d", day),
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			solution.Run(cmd.Context(), *inputFiles[day])
		},
	}

	inputFiles[day] = new(string)
	dayCmd.Flags().StringVarP(inputFiles[day], "input", "i", fmt.Sprintf("inputs/day%d/input", day), "Path to input file for this solution")

	return dayCmd
}

var (
	solutionsList = []solutions.Solution{
		&day1.Solution{},
		&day2.Solution{},
		&day3.Solution{},
		&day4.Solution{},
		&day5.Solution{},
		&day6.Solution{},
		&day7.Solution{},
		&day8.Solution{},
		&day9.Solution{},
		&day10.Solution{},
		&day11.Solution{},
		&day12.Solution{},
		&day13.Solution{},
		&day14.Solution{},
		&day15.Solution{},
		&day16.Solution{},
		&day17.Solution{},
		&day18.Solution{},
		&day19.Solution{},
		&day20.Solution{},
		&day21.Solution{},
		&day22.Solution{},
		&day23.Solution{},
		&day24.Solution{},
		&day25.Solution{},
	}
)

func newDayCommands() map[string]*cobra.Command {
	commands := map[string]*cobra.Command{}
	for i, solution := range solutionsList {
		day := i + 1
		commands[fmt.Sprintf("day%d", day)] = newDayCommand(day, solution)
	}

	return commands
}
