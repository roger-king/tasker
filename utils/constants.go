package utils

import (
	"os"
)

var (
	// DBType -
	DBType = os.Getenv("TASKER_DB_TYPE")

	// DBHost -
	DBHost = os.Getenv("TASKER_DB_HOST")

	// DBUser -
	DBUser = os.Getenv("TASKER_DB_USER")

	// DBPassword -
	DBPassword = os.Getenv("TASKER_DB_PASSWORD")

	// DBPort -
	DBPort = os.Getenv("TASKER_DB_PORT")

	// ScriptRoot -
	ScriptRoot = os.Getenv("TASKER_SCRIPT_ROOT")

	// TASKER_ENV -
	TaskerEnv = os.Getenv("TASKER_ENV")

	// GITHUB_CLIENT_ID
	GithubClientID = os.Getenv("GITHUB_CLIENT_ID")

	// GITHUB_CLIENT_SECRET
	GithubClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")

	// TaskerSecret - used for our signing of access tokens
	TaskerSecret = os.Getenv("TASKER_SECRET")

	// TaskerJWTSecret - used for our signing of jwt tokens
	TaskerJWTSecret = os.Getenv("TASKER_JWT_SECRET")
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
