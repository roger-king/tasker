package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	db "upper.io/db.v3"
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

type Task struct {
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

type TaskService struct {
	Client *redis.Client
	DB     db.Database
}

func (t *TaskService) List() ([]Task, error) {
	var tasks []Task

	if t.Client != nil {
		results, _, err := t.Client.Scan(0, "*", 10).Result()

		if err != nil {
			log.Panic(err)
			return nil, err
		}
		fmt.Println(results)
		// for _, r := range results {
		// 	var task Task
		// 	err := json.Unmarshal([]byte(r), &task)

		// 	if err != nil {
		// 		// log.Error("Failed to marshal data")
		// 		break
		// 	}

		// 	tasks = append(tasks, task)
		// }
	}

	return tasks, nil
}

func (t *TaskService) Create(task *NewInputTask) (*Task, error) {
	fmt.Println(task.Name)
	createdTask := &Task{
		TaskID:    uuid.New().String(),
		Name:      task.Name,
		Schedule:  task.Schedule,
		IsSet:     true,
		Enabled:   true,
		Complete:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if t.Client != nil {
		bt, err := json.Marshal(createdTask)
		if err != nil {
			return nil, err
		}

		err = t.Client.Set(createdTask.TaskID, bt, 0).Err()

		if err != nil {
			return nil, err
		}
	}

	return createdTask, nil
}

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
