package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update [task]",
	Short:   "Update the status of a task",
	Long:    `Update the status of a task in your todo list`,
	Aliases: []string{"edit", "revise"},
	Example: "  nextbox update 2 --status",
	Run:     updateRun,
}

func updateRun(cmd *cobra.Command, args []string) {
	fmt.Println("update")
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
