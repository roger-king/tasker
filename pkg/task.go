package pkg

type ITask interface {
	BeforeCreate()
	Create(t *NewInputTask) (interface{}, error)
	FindOne(taskId string) (Task, error)
	List() ([]Task, error)
	ListEnabledTasks(opts *TaskSearchOptions) ([]Task, error)
	Delete(taskId string) error
	Disable(taskId string) error
}

type Task ITask

// NewInputTask - object to store all parameters for creating a new task
type NewInputTask struct {
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
	Executor string `json:"executor"`
}

// TaskSearchOptions -
type TaskSearchOptions struct {
	Enabled bool `json:"enabled"`
}
