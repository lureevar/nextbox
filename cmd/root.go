package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nextbox",
	Short: "A simple todo list app for the CLI",
	Long:  `A simple todo list application for the CLI based on the CRUD principle`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("verbose", "v", true, "versbose output")
	rootCmd.Flags().BoolP("quiet", "q", false, "quiet output")
}
