package cmd

import (
	"fmt"

	"github.com/segwin/adventofcode-2020/internal/solutions"
	"github.com/segwin/adventofcode-2020/internal/solutions/day1"
	"github.com/segwin/adventofcode-2020/internal/solutions/day2"
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

func newDayCommands() map[string]*cobra.Command {
	return map[string]*cobra.Command{
		"day1": newDayCommand(1, &day1.Solution{}),
		"day2": newDayCommand(2, &day2.Solution{}),
		"day3": newDayCommand(3, &day3.Solution{}),
		"day4": newDayCommand(4, &day4.Solution{}),
		"day5": newDayCommand(5, &day5.Solution{}),
		"day6": newDayCommand(6, &day6.Solution{}),
		"day7": newDayCommand(7, &day7.Solution{}),
		"day8": newDayCommand(8, &day8.Solution{}),
		"day9": newDayCommand(9, &day9.Solution{}),
		//"day10": newDayCommand(10, &day10.Solution{}),
		//"day11": newDayCommand(11, &day11.Solution{}),
		//"day12": newDayCommand(12, &day12.Solution{}),
		//"day13": newDayCommand(13, &day13.Solution{}),
		//"day14": newDayCommand(14, &day14.Solution{}),
		//"day15": newDayCommand(15, &day15.Solution{}),
		//"day16": newDayCommand(16, &day16.Solution{}),
		//"day17": newDayCommand(17, &day17.Solution{}),
		//"day18": newDayCommand(18, &day18.Solution{}),
		//"day19": newDayCommand(19, &day19.Solution{}),
		//"day20": newDayCommand(20, &day20.Solution{}),
		//"day21": newDayCommand(21, &day21.Solution{}),
		//"day22": newDayCommand(22, &day22.Solution{}),
		//"day23": newDayCommand(23, &day23.Solution{}),
		//"day24": newDayCommand(24, &day24.Solution{}),
		//"day25": newDayCommand(25, &day25.Solution{}),
	}
}
