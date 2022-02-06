package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:     "read",
	Short:   "Read your todo list",
	Long:    `Read your todo list and show all the tasks`,
	Aliases: []string{"show"},
	Example: "  nextbox read",
	Run:     readRun,
}

func readRun(cmd *cobra.Command, args []string) {
	fmt.Println("read")
}

func init() {
	rootCmd.AddCommand(readCmd)
}
