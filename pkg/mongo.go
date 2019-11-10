package pkg

import (
	"context"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// BeforeCreate - hook for creation
func (t *Task) BeforeCreate() {
	t.TaskID = uuid.New().String()
	t.Enabled = true
	t.Complete = false
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

// MongoService - a service to interact with a task
type MongoService struct {
	Collection *mongo.Collection
}

// NewMongoService - initializes task service
func NewMongoService(db *mongo.Client) *MongoService {
	collection := db.Database("tasker").Collection("tasks")

	return &MongoService{
		Collection: collection,
	}
}

// List - List operation for task service
func (m *MongoService) List() ([]*Task, error) {
	var tasks []*Task
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := m.Collection.Find(ctx, bson.M{})

	if err != nil {
		log.Info("This is dumb", err)
		return nil, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var task *Task

		err := cur.Decode(&task)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// ListEnabledTasks - List all tasks that are enabled
// TODO: dynamically add the filter
// func (m *MongoService) ListEnabledTasks(opts *TaskSearchOptions) ([]*Task, error) {
// 	var tasks []*Task

// 	results := m.Collection.Find("enabled", opts.Enabled)
// 	err := results.All(&tasks)
// 	return tasks, err
// }

// Create - create operation for task service
func (m *MongoService) Create(newTask *NewInputTask) (*Task, error) {
	task := &Task{
		Name:         newTask.Name,
		Schedule:     newTask.Schedule,
		Executor:     newTask.Executor,
		IsRepeatable: newTask.IsRepeatable,
		Args:         newTask.Args,
	}

	task.BeforeCreate()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	_, err := m.Collection.InsertOne(ctx, task)

	defer cancel()
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Update -
func (m *MongoService) Update(updatedTask *Task) (*Task, error) {
	var task *Task
	var docUpdatedTask bson.M
	// var updateDoc bson.DocElem
	updatedTask.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	bUpdatedTask, err := bson.Marshal(updatedTask)

	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(bUpdatedTask, &docUpdatedTask)

	if err != nil {
		log.Error("Logging error")
		return nil, err
	}

	update := bson.D{
		{
			"$set",
			docUpdatedTask,
		},
	}

	err = m.Collection.FindOneAndUpdate(ctx, bson.M{"taskId": updatedTask.TaskID}, update).Decode(&task)

	if err != nil {
		log.Info("Not working")
		return nil, err
	}

	return task, nil
}

// FindOne -
func (m *MongoService) FindOne(taskID string) (*Task, error) {
	var task Task
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := m.Collection.FindOne(ctx, bson.M{"taskId": taskID}).Decode(&task)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

// // Disable -
// func (m *MongoService) Disable(taskID string) error {
// 	var err error
// 	var task Task

// 	res := m.Collection.Find("taskId", taskID)

// 	err = res.One(&task)

// 	if err != nil {
// 		return err
// 	}

// 	task.Enabled = false
// 	err = res.Update(task)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Delete -
// func (m *MongoService) Delete(taskID string) error {
// 	var err error

// 	res := m.Collection.Find("taskId", taskID)
// 	err = res.Delete()

// 	return err
// }
