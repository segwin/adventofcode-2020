package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	inputDirectory string
)

func newAllCommand() *cobra.Command {
	allCmd := &cobra.Command{
		Use:   "all",
		Short: "Run all solutions back-to-back",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			for i, solution := range solutionsList {
				day := i + 1
				inputFile := strings.Join([]string{inputDirectory, fmt.Sprintf("day%d", day), "input"}, "/")

				divider := "----------"
				if day >= 10 {
					divider += "-"
				}

				fmt.Printf("%s\n  Day %d\n%s\n", divider, day, divider)
				solution.Run(cmd.Context(), inputFile)
				fmt.Println()
			}
		},
	}

	allCmd.Flags().StringVarP(&inputDirectory, "input", "i", "inputs", "Path to directory containing all input files, structured as <dir>/day<N>/input")

	return allCmd
}
