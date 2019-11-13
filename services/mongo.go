package services

import (
	"context"
	"time"

	"github.com/roger-king/tasker/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewMongoConnection - creates a MongoConnection instance
func NewMongoConnection(c string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal("Cannot connect to mongo database")
		return nil, err
	}

	log.Info("Connected to database")
	return client, nil
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
func (m *MongoService) List() ([]*models.Task, error) {
	var tasks []*models.Task
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	options := options.Find()
	options.SetSort(bson.D{{"executor", 1}})

	cur, err := m.Collection.Find(ctx, bson.M{}, options)

	if err != nil {
		log.Info("This is dumb", err)
		return nil, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var task *models.Task

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
// func (m *MongoService) ListEnabledTasks(opts *models.TaskSearchOptions) ([]*models.Task, error) {
// 	var tasks []*models.Task

// 	results := m.Collection.Find("enabled", opts.Enabled)
// 	err := results.All(&tasks)
// 	return tasks, err
// }

// Create - create operation for task service
func (m *MongoService) Create(newTask *models.NewInputTask) (*models.Task, error) {
	task := &models.Task{
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

func (m *MongoService) Update(updatedTask *models.Task) error {
	var docUpdatedTask bson.M
	// var updateDoc bson.DocElem
	updatedTask.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	bUpdatedTask, err := bson.Marshal(updatedTask)

	if err != nil {
		return err
	}

	err = bson.Unmarshal(bUpdatedTask, &docUpdatedTask)

	if err != nil {
		log.Error("Logging error")
		return err
	}

	update := bson.D{
		{
			"$set",
			docUpdatedTask,
		},
	}

	_, err = m.Collection.UpdateOne(ctx, bson.M{"taskId": updatedTask.TaskID}, update)

	if err != nil {
		return err
	}

	return nil
}

// FindOne -
func (m *MongoService) FindOne(taskID string) (*models.Task, error) {
	var task models.Task
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := m.Collection.FindOne(ctx, bson.M{"taskId": taskID}).Decode(&task)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

// Delete -
func (m *MongoService) Delete(taskID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := m.Collection.DeleteOne(ctx, bson.M{"taskId": taskID})

	if err != nil {
		return err
	}

	return nil
}
