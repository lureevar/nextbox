package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/lureevar/nextbox/todo"
	"github.com/spf13/cobra"
)

var eraseCmd = &cobra.Command{
	Use:     "erase",
	Short:   "Erase a task",
	Long:    `Erase a task frow your nextbox`,
	Run:     eraseRun,
	Example: "  " + rootCmd.Name() + " erase 1",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{`purge`},
}

func eraseRun(cmd *cobra.Command, args []string) {
	td := todo.NewTodo(conf.Path)

	argi, err := strconv.Atoi(args[0])
	if err != strconv.ErrSyntax {
		log.Fatalln("error: invalid syntax")
	}

	err = td.EraseExistingTask(argi)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	fmt.Println("Task successfully erased from your nextbox")
}

func init() {
	rootCmd.AddCommand(eraseCmd)
}
