package services

import (
	"time"

	cron "github.com/robfig/cron/v3"
	"github.com/roger-king/tasker/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskService - This is a wrapper around Tasker object
type TaskService struct {
	DB        *mongo.Client
	Scheduler *cron.Cron
}

func (t *TaskService) List() ([]*models.Task, error) {
	m := NewMongoService(t.DB)

	tasks, err := m.List()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskService) Create(i *models.NewInputTask) (*models.Task, error) {
	m := NewMongoService(t.DB)
	createdTask, err := m.Create(i)

	if err != nil {
		return nil, err
	}

	entryId, err := t.Scheduler.AddFunc(i.Schedule, t.Runner(m, createdTask))

	if err != nil {
		// Failed to set job
		return nil, err
	}

	createdTask.EntryID = entryId

	err = m.Update(createdTask)

	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (t *TaskService) Find(id string) (*models.Task, error) {
	m := NewMongoService(t.DB)
	task, err := m.FindOne(id)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *TaskService) Disable(id string) error {
	m := NewMongoService(t.DB)
	task, err := m.FindOne(id)

	if err != nil {
		return err
	}

	task.Enabled = false
	task.UpdatedAt = time.Now()

	err = m.Update(task)

	if err != nil {
		return err
	}

	t.Scheduler.Remove(task.EntryID)

	return nil
}

// Delete -
func (t *TaskService) Delete(id string) error {
	m := NewMongoService(t.DB)
	task, err := m.FindOne(id)

	if err != nil {
		return err
	}

	err = m.Delete(id)

	if err != nil {
		return err
	}

	t.Scheduler.Remove(task.EntryID)

	return nil
}

// NewInputTask - object to store all parameters for creating a new task
type NewInputTask struct {
	Name         string                 `json:"name"`
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
