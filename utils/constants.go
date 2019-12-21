package utils

import (
	"os"
)

var (
	// DBMigrate
	DBMigrate = os.Getenv("TASK_DB_MIGRATE")

	// DBConnectionURL -
	DBConnectionURL = os.Getenv("TASKER_DB_CONNECTION_URL")

	// ScriptRoot -
	ScriptRoot = os.Getenv("TASKER_SCRIPT_ROOT")

	// TaskerEnv -
	TaskerEnv = os.Getenv("TASKER_ENV")

	// GithubClientID
	GithubClientID = os.Getenv("GITHUB_CLIENT_ID")

	// GithubClientSecret
	GithubClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")

	// TaskerSecret - used for our signing of access tokens
	TaskerSecret = os.Getenv("TASKER_SECRET")
)

// HTTPError -
type HTTPError string

func (h HTTPError) String() string {
	return string(h)
}

const (
	// ProcessingError -
	ProcessingError HTTPError = "processing_error"
	// RequestError -
	RequestError HTTPError = "request_error"
)

// Github OAuth Tokens for different permissions

// GithubScopeType -
type GithubScopeType string

const (
	// GithubUserScope -
	GithubUserScope GithubScopeType = "user"
	// GithubRepoScope -
	GithubRepoScope GithubScopeType = "repo"
)
