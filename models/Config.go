package models

type MongoConnectionURL string
type DBType string

type TaskerConfig struct {
	Type               DBType             `required:"true"`
	MongoConnectionURL MongoConnectionURL `required:"true"`
}
