package services

import (
	"context"
	"time"

	"github.com/roger-king/tasker/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoService - a service to interact with a task
type UserService struct {
	Collection *mongo.Collection
}

// NewUserService - initializes task service
func NewUserService(db *mongo.Client) *MongoService {
	collection := db.Database("tasker").Collection("users")

	return &MongoService{
		Collection: collection,
	}
}

func (u *UserService) CreateUser(newUser *models.User) (*models.User, error) {
	newUser.BeforeCreate()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	_, err := u.Collection.InsertOne(ctx, newUser)
	defer cancel()

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
