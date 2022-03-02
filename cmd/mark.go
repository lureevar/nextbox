package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/lureevar/nextbox/todo"
	"github.com/spf13/cobra"
)

var markCmd = &cobra.Command{
	Use:     "mark",
	Short:   "Mark a task as done",
	Long:    `Mark a task as done in your nextbox`,
	Run:     markRun,
	Example: "  " + rootCmd.Name() + " mark 1",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{`do`},
}

func markRun(cmd *cobra.Command, args []string) {
	td := todo.NewTodo(conf.Path)

	argi, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln("error: invalid syntax")
	}

	err = td.MarkTaskUndone(argi)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	fmt.Println("Task successfully marked from your nextbox")
}

func init() {
	rootCmd.AddCommand(markCmd)
}
