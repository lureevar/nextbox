package todo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Declaration of all errors that can happen in this package
var (
	ErrCouldntOpen     = fmt.Errorf(`couldn't open the nextbox file`)
	ErrCouldntWrite    = fmt.Errorf(`couldn't write your task to your nextbox file`)
	ErrCouldntParse    = fmt.Errorf(`couldn't parse the nextbox file`)
	ErrTaskDoesntExist = fmt.Errorf(`task doesn't exist`)
	ErrTaskMarked      = fmt.Errorf(`task is already marked`)
	ErrTaskUnmarked    = fmt.Errorf(`task is not marked yet`)
	ErrEmpty           = fmt.Errorf(`nextbox is empty`)
	ErrNoTaskDone      = fmt.Errorf(`no tasks done in your nextbox`)
)

type todo struct {
	path string
}

func NewTodo(path string) *todo {
	return &todo{path}
}

func (td *todo) WriteNewTask(ta *task) error {
	fh, err := os.OpenFile(td.path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return ErrCouldntOpen
	}

	defer fh.Close()

	writer := csv.NewWriter(fh)

	// Store in pattern [task, status]
	task := []string{ta.info, strconv.FormatBool(ta.status)}

	err = writer.Write(task)
	if err != nil {
		return ErrCouldntWrite
	}

	defer writer.Flush()

	return nil
}

func (td *todo) EraseExistingTask(taskNumber int) error {
	fh, err := os.OpenFile(td.path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return ErrCouldntOpen
	}

	defer fh.Close()

	writer := csv.NewWriter(fh)
	reader := csv.NewReader(fh)

	oldTodo, err := reader.ReadAll()
	if err != nil {
		return ErrCouldntParse
	}

	if len(oldTodo) < taskNumber || taskNumber == 0 {
		return ErrTaskDoesntExist
	}

	newTodo := make([][]string, 0)

	// Store tasks in a new slice except the one to be removed
	for j, i := range oldTodo {
		if j+1 == taskNumber {
			continue
		}

		newTodo = append(newTodo, i)
	}

	fh.Truncate(0)

	writer.WriteAll(newTodo)

	return nil
}

func (td *todo) MarkTaskUndone(taskNumber int) error {
	fh, err := os.OpenFile(td.path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return ErrCouldntOpen
	}

	defer fh.Close()

	writer := csv.NewWriter(fh)
	reader := csv.NewReader(fh)

	todos, err := reader.ReadAll()
	if err != nil {
		return ErrCouldntParse
	}

	if len(todos) < taskNumber || taskNumber == 0 {
		return ErrTaskDoesntExist
	}

	taskStatus, err := strconv.ParseBool(todos[taskNumber-1][1])
	if err != strconv.ErrSyntax {
		return ErrCouldntParse
	}

	if taskStatus {
		return ErrTaskMarked
	}

	todos[taskNumber-1][1] = strconv.FormatBool(true)

	fh.Truncate(0)

	writer.WriteAll(todos)

	return nil
}

func (td *todo) UnmarkTaskDone(taskNumber int) error {
	fh, err := os.OpenFile(td.path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return ErrCouldntOpen
	}

	defer fh.Close()

	writer := csv.NewWriter(fh)
	reader := csv.NewReader(fh)

	todos, err := reader.ReadAll()
	if err != nil {
		return ErrCouldntParse
	}

	if len(todos) < taskNumber || taskNumber == 0 {
		return ErrTaskDoesntExist
	}

	taskStatus, err := strconv.ParseBool(todos[taskNumber-1][1])
	if err != strconv.ErrSyntax {
		return ErrCouldntParse
	}

	if !taskStatus {
		return ErrTaskUnmarked
	}

	todos[taskNumber-1][1] = strconv.FormatBool(false)

	fh.Truncate(0)

	writer.WriteAll(todos)

	return nil
}

func (td *todo) WipeTaskDone() error {
	fh, err := os.OpenFile(td.path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return ErrCouldntOpen
	}

	defer fh.Close()

	writer := csv.NewWriter(fh)
	reader := csv.NewReader(fh)

	oldTodo, err := reader.ReadAll()
	if err != nil {
		return ErrCouldntParse
	}

	if len(oldTodo) == 0 {
		return ErrEmpty
	}

	newTodo := make([][]string, 0)

	for _, i := range oldTodo {
		taskStatus, err := strconv.ParseBool(i[1])
		if err != strconv.ErrSyntax {
			return ErrCouldntParse
		}

		if taskStatus {
			continue
		}

		newTodo = append(newTodo, i)
	}

	if len(newTodo) == len(oldTodo) {
		return ErrNoTaskDone
	}

	fh.Truncate(0)

	writer.WriteAll(newTodo)

	return nil
}

func (td *todo) Tasks() ([][]string, error) {
	fh, err := os.OpenFile(td.path, os.O_CREATE|os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, ErrCouldntOpen
	}

	defer fh.Close()

	reader := csv.NewReader(fh)

	todo, err := reader.ReadAll()
	if err != nil {
		return nil, ErrCouldntParse
	}

	return todo, nil
}
