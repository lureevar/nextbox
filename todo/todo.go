package todo

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

// Declaration of all errors that can happen in this package
const (
	TODO_ERR_COULDNT_OPEN           = `couldn't open the nextbox file`
	TODO_ERR_COULDNT_WRITE          = `couldn't write your task to your nextbox file`
	TODO_ERR_COULDNT_PARSE          = `couldn't parse the nextbox file`
	TODO_ERR_TASK_DOESNT_EXIST      = `task doesn't exist`
	TODO_ERR_COULDNT_PARSE_OR_WRITE = `couldn't parse and\or write to the nextbox file`
	TODO_ERR_TASK_ALREADY_MARKED    = `task is already marked`
	TODO_ERR_TASK_NOT_MARKED        = `task is not marked yet`
	TODO_ERR_EMPTY                  = `nextbox is empty`
	TODO_ERR_NO_TASKS_DONE          = `no tasks done in your nextbox`
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
		return errors.New(TODO_ERR_COULDNT_OPEN)
	}

	defer fh.Close()

	writer := csv.NewWriter(fh)

	// Store in pattern [task, status]
	task := []string{ta.info, strconv.FormatBool(ta.status)}

	err = writer.Write(task)
	if err != nil {
		return errors.New(TODO_ERR_COULDNT_WRITE)
	}

	defer writer.Flush()

	return nil
}

func (td *todo) EraseExistingTask(taskNumber int) error {
	fh, err := os.OpenFile(td.path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return errors.New(TODO_ERR_COULDNT_OPEN)
	}

	defer fh.Close()

	writer := csv.NewWriter(fh)
	reader := csv.NewReader(fh)

	oldTodo, err := reader.ReadAll()
	if err != nil {
		return errors.New(TODO_ERR_COULDNT_PARSE)
	}

	if len(oldTodo) < taskNumber || taskNumber == 0 {
		return errors.New(TODO_ERR_TASK_DOESNT_EXIST)
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

	err = writer.WriteAll(newTodo)
	if err != nil {
		return errors.New(TODO_ERR_COULDNT_PARSE_OR_WRITE)
	}

	return nil
}

func (td *todo) MarkTaskUndone(taskNumber int) error {
	fh, err := os.OpenFile(td.path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return errors.New(TODO_ERR_COULDNT_OPEN)
	}

	defer fh.Close()

	writer := csv.NewWriter(fh)
	reader := csv.NewReader(fh)

	todos, err := reader.ReadAll()
	if err != nil {
		return errors.New(TODO_ERR_COULDNT_PARSE)
	}

	if len(todos) < taskNumber || taskNumber == 0 {
		return errors.New(TODO_ERR_TASK_DOESNT_EXIST)
	}

	taskStatus, err := strconv.ParseBool(todos[taskNumber-1][1])
	if err != nil {
		return errors.New(TODO_ERR_COULDNT_PARSE)
	}

	if taskStatus {
		return errors.New(TODO_ERR_TASK_ALREADY_MARKED)
	}

	todos[taskNumber-1][1] = strconv.FormatBool(true)

	fh.Truncate(0)

	writer.WriteAll(todos)

	return nil
}

func (td *todo) UnmarkTaskDone(taskNumber int) error {
	fh, err := os.OpenFile(td.path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return errors.New(TODO_ERR_COULDNT_OPEN)
	}

	defer fh.Close()

	writer := csv.NewWriter(fh)
	reader := csv.NewReader(fh)

	todos, err := reader.ReadAll()
	if err != nil {
		return errors.New(TODO_ERR_COULDNT_PARSE)
	}

	if len(todos) < taskNumber || taskNumber == 0 {
		return errors.New(TODO_ERR_TASK_DOESNT_EXIST)
	}

	taskStatus, err := strconv.ParseBool(todos[taskNumber-1][1])
	if err != nil {
		return errors.New(TODO_ERR_COULDNT_PARSE)
	}

	if !taskStatus {
		return errors.New(TODO_ERR_TASK_NOT_MARKED)
	}

	todos[taskNumber-1][1] = strconv.FormatBool(false)

	fh.Truncate(0)

	writer.WriteAll(todos)

	return nil
}

func (td *todo) WipeTaskDone() error {
	fh, err := os.OpenFile(td.path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return errors.New(TODO_ERR_COULDNT_OPEN)
	}

	defer fh.Close()

	writer := csv.NewWriter(fh)
	reader := csv.NewReader(fh)

	oldTodo, err := reader.ReadAll()
	if err != nil {
		return errors.New(TODO_ERR_COULDNT_PARSE)
	}

	if len(oldTodo) == 0 {
		return errors.New(TODO_ERR_EMPTY)
	}

	newTodo := make([][]string, 0)

	for _, i := range oldTodo {
		taskStatus, err := strconv.ParseBool(i[1])
		if err != nil {
			return errors.New(TODO_ERR_COULDNT_PARSE)
		}

		if taskStatus {
			continue
		}

		newTodo = append(newTodo, i)
	}

	if len(newTodo) == len(oldTodo) {
		return errors.New(TODO_ERR_NO_TASKS_DONE)
	}

	fh.Truncate(0)

	writer.WriteAll(newTodo)

	return nil
}

func (td *todo) GetTasks() ([][]string, error) {
	fh, err := os.OpenFile(td.path, os.O_CREATE|os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, errors.New(TODO_ERR_COULDNT_OPEN)
	}

	defer fh.Close()

	reader := csv.NewReader(fh)

	todo, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New(TODO_ERR_COULDNT_PARSE)
	}

	return todo, nil
}
