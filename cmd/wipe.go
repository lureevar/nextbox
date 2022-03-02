package cmd

import (
	"fmt"
	"log"

	"github.com/lureevar/nextbox/todo"
	"github.com/spf13/cobra"
)

var wipeCmd = &cobra.Command{
	Use:     "wipe",
	Short:   "Wipe all the marked tasks",
	Long:    `Wipe all the marked tasks in your nextbox`,
	Run:     wipeRun,
	Example: "  " + rootCmd.Name() + " wipe",
	Args:    cobra.NoArgs,
	Aliases: []string{`clean`},
}

func wipeRun(cmd *cobra.Command, args []string) {
	td := todo.NewTodo(conf.Path)

	err := td.WipeTaskDone()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	fmt.Println("Task already done in nextbox wiped successfully")
}

func init() {
	rootCmd.AddCommand(wipeCmd)
}
