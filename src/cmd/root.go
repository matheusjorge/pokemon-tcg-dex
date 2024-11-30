package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func AddCommands(commands []cobra.Command) {
	for _, c := range commands {
		rootCmd.AddCommand(&c)
	}
}
