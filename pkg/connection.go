package pkg

import (
	"log"

	"github.com/go-redis/redis"
	db "upper.io/db.v3"
	"upper.io/db.v3/mongo"
)

type ConnectionType string

const (
	REDIS ConnectionType = "redis"
	MONGO ConnectionType = "mongo"
)

type ConnectionDetails struct {
	Host     string
	User     string
	Password string
	DB       DBName
}

type DBName interface{}

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

// NewRedisConnection -
func NewRedisConnection(d *ConnectionDetails) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     d.Host,
		Password: d.Password, // no password set
		DB:       d.DB.(int), // use default DB
	})

	pong, err := client.Ping().Result()

	if err == nil {
		log.Print("Connected to Redis.", pong)
	}

	return client, err
}
