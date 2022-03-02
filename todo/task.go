package todo

type task struct {
	info   string
	status bool
}

func NewTask(info string, status bool) *task {
	return &task{info, status}
}
