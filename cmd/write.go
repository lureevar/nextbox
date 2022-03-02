package cmd

import (
	"fmt"
	"log"

	"github.com/lureevar/nextbox/todo"
	"github.com/spf13/cobra"
)

var writeCmd = &cobra.Command{
	Use:     "write",
	Short:   "Write down a new task",
	Long:    `Write donw a new task to nextbox`,
	Run:     writeRun,
	Example: "  " + rootCmd.Name() + " add 'Buy some milk'",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{`insert`, `create`, `add`},
}

func writeRun(cmd *cobra.Command, args []string) {
	td := todo.NewTodo(conf.Path)
	ta := todo.NewTask(args[0], false)

	err := td.WriteNewTask(ta)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	fmt.Println("Task successfully written in your nextbox")
}

func init() {
	rootCmd.AddCommand(writeCmd)
}
