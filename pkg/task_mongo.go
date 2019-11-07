package pkg

import (
	"time"

	"github.com/google/uuid"
	"github.com/robfig/cron"
	db "upper.io/db.v3"
)

// Task - Task is the main object describing the collection
type MongoTask struct {
	TaskID    string            `json:"taskId" bson:"taskId"`
	Name      string            `json:"name" bson:"name"`
	Schedule  string            `json:"schedule" bson:"schedule"`
	IsSet     bool              `json:"isSet" bson:"isSet"`
	Enabled   bool              `json:"enabled" bson:"enabled"`
	Complete  bool              `json:"complete" bson:"complete"`
	Args      map[string]string `json:"args" bson:"args"`
	CreatedAt time.Time         `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt" bson:"updatedAt"`
	DeletedAt time.Time         `json:"deletedAt" bson:"deletedAt"`
}

// BeforeCreate - hook for creation
func (t *MongoTask) BeforeCreate() {
	t.TaskID = uuid.New().String()
	t.IsSet = true
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
func (t *TaskService) Create(newTask *NewInputTask, scheduler *cron.Cron) (interface{}, error) {
	task := &MongoTask{
		Name:     newTask.Name,
		Schedule: newTask.Schedule,
	}

	task.BeforeCreate()

	createdTask, err := t.Collection.Insert(task)
	return createdTask, err
}

// FindOne -
func (t *TaskService) FindOne(taskID string) (Task, error) {
	var err error
	var task Task

	res := t.Collection.Find("taskId", taskID)
	err = res.One(&task)

	return task, err
}

// Disable -
func (t *TaskService) Disable(taskID string) error {
	var err error
	var task MongoTask

	res := t.Collection.Find("taskId", taskID)

	err = res.One(&task)

	if err != nil {
		return err
	}

	task.Enabled = false
	err = res.Update(task)

	if err != nil {
		return err
	}

	return nil
}

// Delete -
func (t *TaskService) Delete(taskID string) error {
	var err error

	res := t.Collection.Find("taskId", taskID)
	err = res.Delete()

	return err
}
