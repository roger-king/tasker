package pkg

import "time"

type TaskRedis struct {
	TaskID    string            `bson:"taskId"`
	Name      string            `bson:"name"`
	Schedule  string            `bson:"schedule"`
	IsSet     bool              `bson:"isSet"`
	Enabled   bool              `bson:"enabled"`
	Complete  bool              `bson:"complete"`
	Executor  string            `bson:"executor"`
	Args      map[string]string `bson:"args"`
	CreatedAt time.Time         `bson:"createdAt"`
	UpdatedAt time.Time         `bson:"updatedAt"`
	DeletedAt time.Time         `bson:"deletedAt"`
}
