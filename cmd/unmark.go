package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/lureevar/nextbox/todo"
	"github.com/spf13/cobra"
)

var unmarkCmd = &cobra.Command{
	Use:     "unmark",
	Short:   "Unmark a task that has been done",
	Long:    `Unmark a task that has been done in your nextbox`,
	Run:     unmarkRun,
	Example: "  " + rootCmd.Name() + " unmark 1",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{`undo`},
}

func unmarkRun(cmd *cobra.Command, args []string) {
	td := todo.NewTodo(conf.Path)

	argi, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln("error: invalid syntax")
	}

	err = td.UnmarkTaskDone(argi)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	fmt.Println("Task successfully unmarked from your nextbox")
}

func init() {
	rootCmd.AddCommand(unmarkCmd)
}
