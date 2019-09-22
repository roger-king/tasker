package pkg

import (
	"log"
	"time"

	"github.com/google/uuid"
	db "upper.io/db.v3"
)

// Task - Task is the main object describing the collection
type Task struct {
	TaskID    string    `bson:"taskId"`
	Name      string    `bson:"name"`
	Schedule  string    `bson:"schedule"`
	Enabled   bool      `bson:"enabled"`
	Complete  bool      `bson:"complete"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
	DeletedAt time.Time `bson:"deletedAt"`
}

// NewInputTask - object to store all parameters for creating a new task
type NewInputTask struct {
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
}

// TaskSearchOptions -
type TaskSearchOptions struct {
	Enabled bool `json:"enabled"`
}

// BeforeCreate - hook for creation
func (t *Task) BeforeCreate() {
	t.TaskID = uuid.New().String()
	t.Enabled = true
	t.Complete = false
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

// TaskService - a service to interact with a task
type TaskService struct {
	Collection db.Collection
}

// NewTaskService - initializes task service
func NewTaskService(db db.Database) *TaskService {
	collection := db.Collection("tasks")

	return &TaskService{
		Collection: collection,
	}
}

// List - List operation for task service
func (t *TaskService) List() ([]Task, error) {
	var tasks []Task

	results := t.Collection.Find()
	err := results.All(&tasks)
	return tasks, err
}

// ListEnabledTasks - List all tasks that are enabled
// TODO: dynamically add the filter
func (t *TaskService) ListEnabledTasks(opts *TaskSearchOptions) ([]Task, error) {
	var tasks []Task

	results := t.Collection.Find("enabled", opts.Enabled)
	err := results.All(&tasks)
	return tasks, err
}

// Create - create operation for task service
func (t *TaskService) Create(newTask *NewInputTask) (interface{}, error) {
	task := &Task{
		Name:     newTask.Name,
		Schedule: newTask.Schedule,
	}

	task.BeforeCreate()

	createdTask, err := t.Collection.Insert(task)
	return createdTask, err
}

// Disable -
func (t *TaskService) Disable(taskID string) (bool, error) {
	var err error
	var task Task

	res := t.Collection.Find("taskId", taskID)
	log.Print(res)
	err = res.One(&task)

	if err != nil {
		return false, err
	}

	task.Enabled = false
	err = res.Update(task)

	if err != nil {
		return false, err
	}

	return true, nil
}
