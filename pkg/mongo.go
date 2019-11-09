package pkg

import (
	"time"

	"github.com/google/uuid"
	"github.com/robfig/cron"
	db "upper.io/db.v3"
)

// Task - Task is the main object describing the collection
type MongoTask struct {
	TaskID       string                 `json:"taskId" bson:"taskId"`
	EntryID      int                    `json:"entryId" bson:"entryId"`
	Name         string                 `json:"name" bson:"name"`
	Executor     string                 `json:"executor" bson:"executor"`
	Schedule     string                 `json:"schedule" bson:"schedule"`
	IsRepeatable bool                   `json:"isRepeatable" bson:"isRepeatable"`
	IsSet        bool                   `json:"isSet" bson:"isSet"`
	Enabled      bool                   `json:"enabled" bson:"enabled"`
	Complete     bool                   `json:"complete" bson:"complete"`
	Args         map[string]interface{} `json:"args" bson:"args"`
	CreatedAt    time.Time              `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time              `json:"updatedAt" bson:"updatedAt"`
	DeletedAt    time.Time              `json:"deletedAt" bson:"deletedAt"`
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

// MongoService - a service to interact with a task
type MongoService struct {
	Collection db.Collection
}

// NewMongoService - initializes task service
func NewMongoService(db db.Database) *MongoService {
	collection := db.Collection("tasks")

	return &MongoService{
		Collection: collection,
	}
}

// List - List operation for task service
func (m *MongoService) List() ([]*Task, error) {
	var tasks []*Task

	results := m.Collection.Find()
	err := results.All(&tasks)
	return tasks, err
}

// ListEnabledTasks - List all tasks that are enabled
// TODO: dynamically add the filter
func (m *MongoService) ListEnabledTasks(opts *TaskSearchOptions) ([]*Task, error) {
	var tasks []*Task

	results := m.Collection.Find("enabled", opts.Enabled)
	err := results.All(&tasks)
	return tasks, err
}

// Create - create operation for task service
func (m *MongoService) Create(newTask *NewInputTask, scheduler *cron.Cron) (interface{}, error) {
	task := &MongoTask{
		Name:     newTask.Name,
		Schedule: newTask.Schedule,
	}

	task.BeforeCreate()

	createdTask, err := m.Collection.Insert(task)

	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

// FindOne -
func (m *MongoService) FindOne(taskID string) (*Task, error) {
	var err error
	var task Task

	res := m.Collection.Find("taskId", taskID)
	err = res.One(&task)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

// Disable -
func (m *MongoService) Disable(taskID string) error {
	var err error
	var task MongoTask

	res := m.Collection.Find("taskId", taskID)

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
func (m *MongoService) Delete(taskID string) error {
	var err error

	res := m.Collection.Find("taskId", taskID)
	err = res.Delete()

	return err
}
