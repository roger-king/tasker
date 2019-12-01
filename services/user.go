package services

import (
	"context"
	"time"

	"github.com/roger-king/tasker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoService - a service to interact with a task
type UserService struct {
	Collection *mongo.Collection
}

// NewUserService - initializes task service
func NewUserService(db *mongo.Client) *UserService {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	collection := db.Database("tasker").Collection("users")
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{Keys: bson.M{
		"username": 1,
	}, Options: options.Index().SetUnique(true)})

	return &UserService{
		Collection: collection,
	}
}

func (u *UserService) CreateUser(newUser *models.User) (*models.User, error) {
	newUser.BeforeCreate()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := u.Collection.FindOneAndReplace(ctx, bson.M{"username": newUser.UserName}, newUser)
	if result.Err() != nil {
		_, err := u.Collection.InsertOne(ctx, newUser)

		if err != nil {
			return nil, err
		}
	}

	return newUser, nil
}

func (u *UserService) FindUser(username string) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := u.Collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
