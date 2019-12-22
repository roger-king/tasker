package services

import "github.com/google/wire"

// var ServiceSet = wire.NewSet(NewGithubAuthService, NewUserService, NewSettingService, NewTaskService, NewMongoService, NewMongoConnection)

// ServiceSet -
var ServiceSet = wire.NewSet(NewGithubAuthService, NewUserService, NewDBConnection)
