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

func (s *SettingService) ListPluginSettings(filters map[string]interface{}) ([]*models.PluginSetting, error) {
	var settings []*models.PluginSetting
	var docFilters bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(bson.D{{"repo_name", 1}})

	if _, ok := filters["skip"]; ok {
		if filters["skip"].(int64) > 0 {
			options.SetSkip(filters["skip"].(int64))
			delete(filters, "skip")
		}
	}

	bFilters, err := bson.Marshal(&filters)

	if err != nil {
		return nil, err
	}

	if err := bson.Unmarshal(bFilters, &docFilters); err != nil {
		return nil, err
	}

	cur, err := s.Collection.Find(ctx, docFilters, options)

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

func (s *SettingService) FindPluginSettingByRepo(repoName string) (*models.PluginSetting, error) {
	var pluginSetting models.PluginSetting
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := s.Collection.FindOne(ctx, bson.M{"repo_name": repoName}).Decode(&pluginSetting)

	if err != nil {
		return nil, err
	}

	return &pluginSetting, nil
}

func (s *SettingService) ToggleActiveSettingPluginRepo(toggleInput *models.ToggleActiveSetting) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := s.Collection.UpdateOne(ctx, bson.M{"repo_name": toggleInput.RepoName}, bson.M{"$set": bson.M{"active": toggleInput.Active}})

	if err != nil {
		return err
	}

	return nil
}
