package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/roger-king/tasker/models"
	log "github.com/sirupsen/logrus"
)

// NewMongoConnection - creates a MongoConnection instance
func NewMongoConnection(tc *models.TaskerConfig) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(string(tc.MongoConnectionURL)))

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
