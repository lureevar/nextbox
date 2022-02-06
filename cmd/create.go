package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new task",
	Long:    `Create a new task and put that in your todo list`,
	Example: `  nextbox create 'Buy some milk'`,
	Aliases: []string{"add", "insert"},
	Run:     createRun,
}

func createRun(cmd *cobra.Command, args []string) {
	fmt.Println("create")
}

func init() {
	rootCmd.AddCommand(createCmd)
}
