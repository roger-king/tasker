package pkg

import (
	"fmt"
	"plugin"
	"time"

	cron "github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITask interface {
	BeforeCreate()
	Create(t *NewInputTask) (interface{}, error)
	FindOne(taskId string) (Task, error)
	List() ([]Task, error)
	ListEnabledTasks(opts *TaskSearchOptions) ([]Task, error)
	Delete(taskId string) error
	Disable(taskId string) error
}

type Payload interface {
	Run(args map[string]interface{}) error
}

type Task struct {
	TaskID       string                 `json:"taskId" bson:"taskId"`
	EntryID      cron.EntryID           `json:"entryId" bson:"entryId"`
	Name         string                 `json:"name" bson:"name"`
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

// TaskService - This is a wrapper around Tasker object
type TaskService struct {
	DB        *mongo.Client
	Scheduler *cron.Cron
}

func (t *TaskService) List() ([]*Task, error) {
	m := NewMongoService(t.DB)

	tasks, err := m.List()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskService) Create(i *NewInputTask) (*Task, error) {
	m := NewMongoService(t.DB)
	createdTask, err := m.Create(i)

	if err != nil {
		return nil, err
	}

	entryId, err := t.Scheduler.AddFunc(i.Schedule, func() {
		this, err := m.FindOne(createdTask.TaskID)

		if err != nil {
			return
		}

		if this.Enabled && !this.Complete {
			plug, err := plugin.Open(fmt.Sprintf("./plugins/%s.so", this.Executor))

			if err != nil {
				fmt.Println(err)
				return
			}

			run, err := plug.Lookup("Run")

			if err != nil {
				fmt.Println(err)
				return
			}

			err = run.(func(map[string]interface{}) error)(this.Args)

			if err != nil {
				fmt.Print(err)
				return
			}

			if !this.IsRepeatable {
				this.Complete = true
				this.Enabled = false
				this.UpdatedAt = time.Now()
				this.DeletedAt = time.Now()

				err = m.Update(this)

				if err != nil {
					log.Error("Failed to mark as complete")
					return
				}

				t.Scheduler.Remove(cron.EntryID(this.EntryID))
			}
		}
	})

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

func (t *TaskService) Find(id string) (*Task, error) {
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
