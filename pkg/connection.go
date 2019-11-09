package pkg

import (
	"log"

	db "upper.io/db.v3"
	"upper.io/db.v3/mongo"
)

type ConnectionDetails struct {
	Host     string `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	DBName   string `required:"true"`
}

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

// NewMongoConnection - creates a MongoConnection instance
func NewMongoConnection() (db.Database, error) {
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

	return session, nil
}
