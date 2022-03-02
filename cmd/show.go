package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lureevar/nextbox/todo"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:     "show",
	Short:   "Show all tasks",
	Long:    `Show all tasks in nextbox`,
	Run:     showRun,
	Example: "  " + rootCmd.Name() + " show",
	Args:    cobra.NoArgs,
}

func showRun(cmd *cobra.Command, args []string) {
	td := todo.NewTodo(conf.Path)

	todo, err := td.GetTasks()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	if len(todo) == 0 {
		fmt.Println("Nothing to do here")
		os.Exit(0)
	}

	fmt.Println("Here is your todo list:")

	for j, i := range todo {
		taskStatus, _ := strconv.ParseBool(i[1])

		fmt.Printf("%d - ", j+1)

		if taskStatus {
			fmt.Printf("\033[9m%s\033[0m\n", i[0])
		} else {
			fmt.Println(i[0])
		}
	}
}

func init() {
	rootCmd.AddCommand(showCmd)
}
