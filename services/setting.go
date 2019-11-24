package services

import (
	"context"
	"time"

	"github.com/roger-king/tasker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SettingService - a service to interact with the settings of tasker
type SettingService struct {
	Collection *mongo.Collection
}

// NewSettingService - initializes task service
func NewSettingService(db *mongo.Client) *SettingService {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	collection := db.Database("tasker").Collection("settings")
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{Keys: bson.M{
		"repo_name": 1,
	}, Options: options.Index().SetUnique(true)})

	return &SettingService{
		Collection: collection,
	}
}

func (s *SettingService) CreatePluginSetting(input *models.PluginSetting) (*models.PluginSetting, error) {
	input.BeforeCreate()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	_, err := s.Collection.InsertOne(ctx, input)

	defer cancel()
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (s *SettingService) ListPluginSettings() ([]*models.PluginSetting, error) {
	var settings []*models.PluginSetting
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	options := options.Find()
	options.SetSort(bson.D{{"executor", 1}})

	cur, err := s.Collection.Find(ctx, bson.M{}, options)

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var setting models.PluginSetting

		err := cur.Decode(&setting)

		if err != nil {
			return nil, err
		}

		settings = append(settings, &setting)
	}

	return settings, nil
}
