package models

import (
	"time"

	"github.com/google/uuid"
	cron "github.com/robfig/cron/v3"
)

// NewInputTask - object to store all parameters for creating a new task
type NewInputTask struct {
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	Args         map[string]interface{} `json:"args"`
	Schedule     string                 `json:"schedule"`
	IsRepeatable bool                   `json:"isRepeatable"`
	Executor     string                 `json:"executor"`
	EntryID      cron.EntryID           `json:"entryId"`
}

// TaskSearchOptions -
type TaskSearchOptions struct {
	Enabled bool `json:"enabled"`
}

type ITask interface {
	BeforeCreate()
	Create(t *NewInputTask) (interface{}, error)
	FindOne(taskId string) (Task, error)
	List() ([]Task, error)
	ListEnabledTasks(opts *TaskSearchOptions) ([]Task, error)
	Delete(taskId string) error
	Disable(taskId string) error
}

type Task struct {
	TaskID       string                 `json:"taskId" bson:"taskId"`
	EntryID      cron.EntryID           `json:"entryId" bson:"entryId"`
	Name         string                 `json:"name" bson:"name"`
	Description  string                 `json:"description" bson:"description"`
	Executor     string                 `json:"executor" bson:"executor"`
	Schedule     string                 `json:"schedule" bson:"schedule"`
	IsRepeatable bool                   `json:"isRepeatable" bson:"isRepeatable"`
	Enabled      bool                   `json:"enabled" bson:"enabled"`
	Complete     bool                   `json:"complete" bson:"complete"`
	Args         map[string]interface{} `json:"args" bson:"args"`
	CreatedAt    time.Time              `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time              `json:"updatedAt" bson:"updatedAt"`
	DeletedAt    time.Time              `json:"deletedAt" bson:"deletedAt"`
}

// BeforeCreate - hook for creation
func (t *Task) BeforeCreate() {
	t.TaskID = uuid.New().String()
	t.Enabled = true
	t.Complete = false
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}
