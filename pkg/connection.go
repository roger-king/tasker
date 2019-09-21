package pkg

import (
	"log"

	db "upper.io/db.v3"
	"upper.io/db.v3/mongo"
)

// MongoConnectionOptions -
type MongoConnectionOptions struct {
	AuthSource string
}

// MongoConnectionURL -
type MongoConnectionURL struct {
	DBName   string
	Host     string
	Port     string
	User     string
	Password string
	Options  MongoConnectionOptions
}

// MongoConnection - object to store connection to mongo db
type MongoConnection struct {
	DB db.Database
}

// NewMongoConnection - creates a MongoConnection instance
func NewMongoConnection() (*MongoConnection, error) {
	options := make(map[string]string)
	options["authSource"] = "admin"

	settings := mongo.ConnectionURL{
		Database: `cronus`,
		Host:     `127.0.0.1`,
		User:     `appuser`,
		Password: `appuser`,
		Options:  options,
	}

	session, err := mongo.Open(settings)

	if err != nil {
		log.Fatal(err)
	}

	return &MongoConnection{
		DB: session,
	}, nil
}
