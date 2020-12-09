package cmd

import (
	"github.com/spf13/cobra"
)

func New(name string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   name,
		Short: "Collection of solutions for the Advent of Code 2020 event",
	}

	for _, cmd := range newDayCommands() {
		rootCmd.AddCommand(cmd)
	}

	return rootCmd
}
