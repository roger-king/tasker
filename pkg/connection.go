package pkg

import (
	db "upper.io/db.v3"
	"upper.io/db.v3/mongo"
)

// MongoConnection - object to store connection to mongo db
type MongoConnection struct {
	DB db.Database
}

// NewMongoConnection - creates a MongoConnection instance
func NewMongoConnection() (*MongoConnection, error) {
	settings := mongo.ConnectionURL{
		Database: "cronus",
		Host:     "localhost",
		User:     "appuser",
		Password: "appuser",
	}

	session, err := mongo.Open(settings)

	if err != nil {
		return nil, err
	}

	return &MongoConnection{
		DB: session,
	}, nil
}
