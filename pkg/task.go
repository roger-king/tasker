package pkg

import (
	"time"

	"github.com/google/uuid"
	db "upper.io/db.v3"
)

// Task - Task is the main object describing the collection
type Task struct {
	TaskID    uuid.UUID `bson:"taskId"`
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

// BeforeCreate - hook for creation
func (t *Task) BeforeCreate() {
	t.TaskID = uuid.New()
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
