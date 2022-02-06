package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete [task]",
	Short:   "Delete an existing task",
	Long:    `Delete an existing taks from your todo list`,
	Aliases: []string{"remove", "purge"},
	Example: "  nextbox delete 2",
	Run:     deleteRun,
}

func deleteRun(cmd *cobra.Command, args []string) {
	fmt.Println("delete")
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
